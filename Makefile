#!make
# You can change file env like .env or .env.*.local
include .env.local

debug:
	go run cmd/server/main.go
.PHONY: debug

build:
	go build -o deploy/engine ./cmd
.PHONY: build

inject:
	wire github.com/MochamadAkbar/go-restful/injector
.PHONY: inject

lint:
	golangci-lint run ./...
.PHONY: lint

migrateinit:
	migrate create -ext sql -dir db/migrations init_schema
.PHONY: migrateinit

migrateup:
	migrate -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}" -path db/migrations --verbose up
.PHONY: migrateup

migratedown:
	migrate -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}" -path db/migrations --verbose down
.PHONY: migratedown

test:
	go test -v ./...
.PHONY: test