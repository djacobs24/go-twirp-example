.PHONY: protos haberdasher

build: protos

protos: haberdasher

clean:
	rm ./rpc/haberdasher/*.go

haberdasher:
	protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. ./rpc/haberdasher/*.proto
