.PHONY: build clean test

GOPATH?=$(HOME)/go

build:
	go build -o $(GOPATH)/bin/go-api-starter ./cmd/main.go 

clean:
	rm -rf $(GOPATH)/bin/go-api-starter

test:
	go test -race ./...