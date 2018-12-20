FROM golang

ADD . /go/src/github.com/kladov/go-api-starter

WORKDIR /go/src/github.com/kladov/go-api-starter

RUN make update-swagger

RUN make build

ENTRYPOINT /go/bin/go-api-starter

EXPOSE 8080