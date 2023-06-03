.DEFAULT_GOAL := run

.PHONY: fmt lint vet build

fmt:
	gofmt -w .

lint: fmt
	golint .

vet: fmt
	go vet .

build: vet
	go build hello.go

run:
	go run .