debug:
	go run cmd/main.go
.PHONY: debug

build:
	go build -o deploy/engine ./cmd
.PHONY: build

test:
	go test -v ./...
.PHONY: test

lint:
	golangci-lint run ./...
.PHONY: lint