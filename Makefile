debug:
	go run cmd/main.go

lint:
	golangci-lint run ./...

.PHONY: debug