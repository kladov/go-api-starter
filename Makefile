.PHONY: build clean
build:
	go build -o $(GOPATH)/bin/go-app-starter ./cmd/main.go 

clean:
	rm -rf $(GOPATH)/bin/go-app-starter