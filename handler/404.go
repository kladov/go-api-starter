package handler

import (
	"encoding/json"
	"net/http"
)

// NotFoundRoute
type NotFoundRoute struct{}

// MatchHandler
func (h *NotFoundRoute) MatchHandler(r *http.Request) http.HandlerFunc {
	return NotFoundHandler
}

var NotFoundHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	json.NewEncoder(w).Encode(Response{"Resource not found"})
}
