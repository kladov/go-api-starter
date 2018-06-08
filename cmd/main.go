package main

import (
	"fmt"
	"net/http"

	"github.com/kladov/go-api-starter/handler"
	"github.com/kladov/go-api-starter/middleware"
	"github.com/kladov/go-api-starter/mux"
)

func main() {
	r := mux.NewSimpleHTTPRouter()
	r.WithRoutes([]mux.Route{
		&handler.IndexRoute{},
		&handler.NotFoundRoute{},
	})
	r.WithMiddleware([]mux.Middleware{
		middleware.ContentTypeHeader,
	})
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Errorf("Error while serve http request: %+s", err)
	}
}
