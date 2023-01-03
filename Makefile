debug:
	go run cmd/main.go
.PHONY: debug

build:
	go build -o deploy/engine ./cmd
.PHONY: build

inject:
	cd injector && wire

lint:
	golangci-lint run ./...
.PHONY: lint

test:
	go test -v ./...
.PHONY: test