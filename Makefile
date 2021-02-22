GRPC_SOURCE ?= api
GRPC_FILES  ?= kvapi.proto
GRPC_OUTPUT ?= grpc/kvapi

GO_BUILD_PATH = bin/app

.PHONY: all
all: grpc build

.PHONY: build
build:
	@go build -o $(GO_BUILD_PATH) cmd/app/main.go

.PHONY: run_docker
run_docker: 
	@docker-compose up -d

.PHONY: build_docker
build_docker: build 
	@docker-compose build

.PHONY: stop_docker
stop_docker:
	@docker-compose stop

.PHONY: generate-grpc
generate-grpc:	
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
		--host 0.0.0.0\
		--port 8081 \
		--path api \
		--proto kvapi.proto


