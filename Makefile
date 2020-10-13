.PHONY: build
build:
	go build -v ./cmd/back/main.go

.DEFAULT_GOAL := build