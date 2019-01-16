.PHONY: build protos clean hat server client

build: protos

protos: hat

clean:
	rm ./rpc/hats/*.go

hat:
	protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. ./rpc/hats/*.proto

server:
	go run ./cmd/server/main.go

client:
	go run ./cmd/client/main.go
	