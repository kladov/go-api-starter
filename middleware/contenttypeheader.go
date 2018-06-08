package middleware

import (
	"net/http"

	"github.com/kladov/go-api-starter/mux"
)

// ContentTypeHeaderMiddleware inject Content-Type header to response
var ContentTypeHeader mux.Middleware = func(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	}
}
