mock-generate:
	@go install github.com/golang/mock/mockgen@latest
	@go generate ./...
	@go mod tidy

proto-generate:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@cd protobuf
	@go generate ./...
	@go mod tidy

run-embed:
	@go run web/web.go