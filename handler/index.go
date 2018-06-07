package handler

import (
	"encoding/json"
	"net/http"

	"github.com/kladov/go-api-starter/mux"
)

// IndexRoute
type IndexRoute struct{}

// MatchHandler
func (h *IndexRoute) MatchHandler(r *http.Request) mux.Handler {
	if r.RequestURI == "/" && r.Method == "GET" {
		return &IndexHandler{}
	}

	return nil
}

// IndexHandler
type IndexHandler struct{}

// Handle
func (h *IndexHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(Response{"Hello world"})
}
