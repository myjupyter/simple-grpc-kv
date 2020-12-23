GRPC_SOURCE=api
GRPC_FILES=kvapi.proto
GRPC_OUTPUT=grpc/kvapi

GO_BUILD_PATH=bin

.PHONY: all
all: grpc build

.PHONY: grpc
grpc:	
	@mkdir -p $(GRPC_OUTPUT) 
	@protoc -I$(GRPC_SOURCE) \
		-I "$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/" \
		--go_out=$(GRPC_OUTPUT) \
		--go-grpc_out=$(GRPC_OUTPUT) \
		--grpc-gateway_out=logtostderr=true:$(GRPC_OUTPUT) $(GRPC_FILES) 

.PHONY: grpc-client
grpc-client:
	@evans repl \
		--path "$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/" \
		-p 8081 \
		--path api \
		--proto kvapi.proto

.PHONY: build
build:
	@go build -o $(GO_BUILD_PATH) cmd/app/main.go
