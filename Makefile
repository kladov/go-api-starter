.PHONY: build clean
build:
	go build -o $(GOPATH)/bin/go-api-starter ./cmd/main.go 

clean:
	rm -rf $(GOPATH)/bin/go-api-starter