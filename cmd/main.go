package main

import (
	"fmt"
	"net/http"

	"github.com/kladov/go-api-starter/handler"
	"github.com/kladov/go-api-starter/middleware"
	"github.com/kladov/go-api-starter/mux"
	"github.com/kladov/go-api-starter/swagger"
)

func main() {
	r := mux.NewSimpleHTTPRouter()
	r.WithRoutes([]mux.Route{
		&handler.IndexRoute{},
		&handler.NotFoundRoute{},
	})
	r.WithMiddleware([]mux.Middleware{
		middleware.ContentTypeHeader,
		middleware.CORS,
	})

	go initSwaggerHTTP()

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Errorf("Error while serve http request: %s", err)
	}
}

func initSwaggerHTTP() {
	swaggerMux := http.NewServeMux()
	swaggerMux.Handle("/", swagger.NewHTTPHandler())
	err := http.ListenAndServe(":8081", swaggerMux)
	if err != nil {
		fmt.Errorf("Error while serve http request: %s", err)
	}
}
