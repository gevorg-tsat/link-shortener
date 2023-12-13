PHONY: generate-structs
generate-structs:
	mkdir -p internal/shortener_v1
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	export PATH="$PATH:$(go env GOPATH)/bin"
	protoc --go_opt=paths=source_relative --go_out=internal/shortener_v1 \
			--go-grpc_out=internal/shortener_v1 --go-grpc_opt=paths=source_relative \
			api/shortener_v1/shortener.proto
	mv internal/shortener_v1/api/shortener_v1/* internal/shortener_v1/
	rm -rf internal/shortener_v1/api

build:
	go build cmd/app/app.go

run-in-memory: build
	./app --storage=in-memory

run-postgres: build
	./app --storage=postgres