package handler

import (
	"encoding/json"
	"net/http"
)

// IndexRoute
type IndexRoute struct{}

// MatchHandler
func (h *IndexRoute) MatchHandler(r *http.Request) http.HandlerFunc {
	if r.RequestURI == "/" && r.Method == "GET" {
		return IndexHandler
	}

	return nil
}

// IndexHandler api root handler
var IndexHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(Response{"Hello world"})
}
