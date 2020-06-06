SERVER_OUT := "build/server"
CLIENT_OUT := "build/client"
SERVER_PKG_BUILD := "./cmd/server"
CLIENT_PKG_BUILD := "./cmd/client"

.PHONY: all api api/api.pb.go build_server build_client clean help

build: build_server build_client

api/api.pb.go: proto/api/api.proto
	@protoc -I proto/api \
		-I${GOPATH}/src \
		-I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.6/third_party/googleapis \
		--go_out=plugins=grpc:api \
		proto/api/api.proto

api/api.pb.gw.go: proto/api/api.proto
	@protoc -I proto/api/ \
		-I${GOPATH}/src \
		-I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.6/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:api \
		proto/api/api.proto

api/api.swagger.json: proto/api/api.proto
	@protoc -I proto/api/ \
		-I${GOPATH}/src \
		-I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.6/third_party/googleapis \
		--swagger_out=logtostderr=true:api \
		proto/api/api.proto

api: api/api.pb.go api/api.pb.gw.go api/api.swagger.json ## Auto-generate grpc go sources

dep: ## Get the dependencies
	@go get -v -d ./...

build_server: dep api ## Build the binary file for server
	@go build -i -v -o $(SERVER_OUT) $(SERVER_PKG_BUILD)

build_client: dep api ## Build the binary file for client
	@go build -i -v -o $(CLIENT_OUT) $(CLIENT_PKG_BUILD)

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
