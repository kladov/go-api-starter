package handler

import (
	"encoding/json"
	"net/http"

	"github.com/kladov/go-api-starter/mux"
)

// NotFoundRoute
type NotFoundRoute struct{}

// MatchHandler
func (h *NotFoundRoute) MatchHandler(r *http.Request) mux.Handler {
	return &NotFoundHandler{}
}

// NotFoundHandler
type NotFoundHandler struct{}

// Handle
func (h *NotFoundHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(Response{"Resource not found"})
}
