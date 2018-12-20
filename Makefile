.PHONY: build clean test

GOPATH?=$(HOME)/go

build:
	go build -o $(GOPATH)/bin/go-api-starter ./cmd/main.go 

clean:
	rm -rf $(GOPATH)/bin/go-api-starter

test:
	go test -race ./...

update-swagger:
	go get -u github.com/sergei-svistunov/gostatic2lib
	rm -rf /tmp/swagger-ui
	cp -r swagger-ui /tmp/swagger-ui
	mkdir -p /tmp/swagger-ui/html
	cp ./swagger/swagger.json /tmp/swagger-ui/html/swagger.json
	cd /tmp/swagger-ui; \
		cat ./dist/index.html | perl -pe 's/https?:\/\/petstore.swagger.io\/v2\///g' > ./html/index.html; \
		cp ./dist/*.js ./html; \
		cp ./dist/*.css ./html; \
		cp ./dist/*.png ./html
	gostatic2lib -out ./swagger/static.go -package swagger -path /tmp/swagger-ui/html
