debug:
	go run cmd/main.go

build:
	go build -o deploy/engine ./cmd

test:
	go test -v ./...

lint:
	golangci-lint run ./...

.PHONY: debug