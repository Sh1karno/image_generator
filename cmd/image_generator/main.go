package main

import (
	"context"
	"github.com/Sh1karno/image_generator/internal/config"
	"github.com/Sh1karno/image_generator/pkg/app"
	"log"
)

func main() {
	ctx := context.Background()
	cfg := config.New()
	a, aErr := app.NewApp(ctx, cfg)
	if aErr != nil {
		log.Fatalln(aErr)
	}

	err := a.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
