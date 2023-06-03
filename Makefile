.DEFAULT_GOAL := build

.PHONY: fmt lint vet

fmt:
	gofmt -w .

lint: fmt
	golint .

vet: fmt
	go vet .

build:
	go build hello.go