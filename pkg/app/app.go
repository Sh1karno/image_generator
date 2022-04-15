package app

import (
	"context"
	"github.com/Sh1karno/image_generator/internal/app/image_generator"
	pb "github.com/Sh1karno/image_generator/pkg/api/github.com/Sh1karno/image_generator"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
)

type IConfig interface {
	GetPort() string
}

type App struct {
	ctx context.Context
	mux *runtime.ServeMux

	Impl       *image_generator.Implementation
	grpcServer *grpc.Server
	cfg        IConfig
}

// NewApp is a constructor
func NewApp(ctx context.Context, cfg IConfig) (*App, error) {
	a := &App{}
	err := a.initDeps(ctx)
	a.ctx = ctx
	a.mux = runtime.NewServeMux()
	a.cfg = cfg

	return a, err
}

// initDeps initialize dependencies
func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initImpl,
		a.initServer,
	}

	for _, fn := range inits {
		if err := fn(ctx); err != nil {
			return err
		}
	}

	return nil
}

// initImpl initialize API impl
func (a *App) initImpl(ctx context.Context) error {
	ImageGeneratorService := image_generator.NewImageGeneratorService()
	a.Impl = ImageGeneratorService

	return nil
}

// initServer ...
func (a *App) initServer(ctx context.Context) error {
	opts := []grpc.ServerOption{}
	a.grpcServer = grpc.NewServer(opts...)
	return nil
}

// Run runs application
func (a *App) Run() error {

	grpcServer := grpc.NewServer()
	pb.RegisterImageGeneratorServer(grpcServer, a.Impl)
	lis, _ := net.Listen("tcp", "localhost:8842")

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	var group errgroup.Group

	group.Go(func() error {
		return grpcServer.Serve(lis)
	})
	group.Go(func() error {
		return pb.RegisterImageGeneratorHandlerFromEndpoint(a.ctx, a.mux, ":8842", opts)
	})
	group.Go(func() error {
		return http.ListenAndServe(":8843", a.mux)
	})

	return group.Wait()
}
