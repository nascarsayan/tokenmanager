all: tidy build

generate:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pkg/token/token.proto

tidy:
	go mod tidy -compat=1.17

build:
	go build -o bin/tokenserver src/server/*.go
	go build -o bin/tokenclient src/client/*.go
