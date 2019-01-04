FROM golang

ADD . /go/src/github.com/kladov/go-api-starter

WORKDIR /go/src/github.com/kladov/go-api-starter

ENV GO111MODULE=on

RUN make update-swagger

RUN make deps

RUN make build

ENTRYPOINT /go/bin/go-api-starter

EXPOSE 8080