debug:
	go run cmd/main.go

test:
	go test -v ./...

lint:
	golangci-lint run ./...

.PHONY: debug