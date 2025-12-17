package main

//go:generate protoc --go_out=. --go-grpc_out=. ./proto/myservice.proto

func main() {
	// go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	// go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

	// go generate ./...
}