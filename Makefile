export PACKAGE_PATH=github.com/Sh1karno/image_generator

LOCAL_BIN:=$(CURDIR)/bin
PROTOGEN_BIN:=$(LOCAL_BIN)/protogen

.PHONY: generate
generate:
	mkdir -p "pkg/api"
	protoc -I /usr/local/include -I . \
        -I ${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
        -I ${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.6.3 \
        -I ${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/v2@v2.7.3 \
        --validate_out=lang=go:./pkg/api \
        --grpc-gateway_out=logtostderr=true:./pkg/api/ \
        --swagger_out=allow_merge=true,merge_file_name=./pkg/api/${PACKAGE_PATH}/image_generator:. \
        --go_out=plugins=grpc:./pkg/api \
        ./api/proto/*.proto
