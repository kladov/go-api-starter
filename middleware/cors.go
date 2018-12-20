package middleware

import (
	"net/http"

	"github.com/kladov/go-api-starter/mux"
)

// ContentTypeHeaderMiddleware inject Content-Type header to response
var CORS mux.Middleware = func(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		addCORSHeader(w)
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	}
}

func addCORSHeader(res http.ResponseWriter) {
	headers := res.Header()
	headers.Add("Access-Control-Allow-Origin", "*")
	headers.Add("Vary", "Origin")
	headers.Add("Vary", "Access-Control-Request-Method")
	headers.Add("Vary", "Access-Control-Request-Headers")
	headers.Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token")
	headers.Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
}
