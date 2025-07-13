.PHONY: deps api worker subscriber

deps:
	go mod download

api:
	go run ./cmd/api

worker:
	go run ./cmd/worker

subscriber:
	go run ./cmd/subscriber
