.PHONY: build run test clean

build:
	@echo "Building binary..."
	go build -o bin/server ./cmd/server/main.go

run:
	@go run ./cmd/server/main.go

test:
	@go test ./...

tidy:
	@go mod tidy