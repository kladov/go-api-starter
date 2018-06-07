package mux

import (
	"net/http"
)

// Route go-between Handler and http.Request
type Route interface {
	// MatchHandler return related Handler based on http.Request (nil if not matched)
	MatchHandler(r *http.Request) Handler
}

// Handler can handle http request
type Handler interface {
	Handle(http.ResponseWriter, *http.Request)
}

// SimpleHTTPRouter simple implementation of HTTP request router
type SimpleHTTPRouter struct {
	routes []Route
}

// NewSimpleHTTPRouter crete new instance of SimpleHTTPRouter
func NewSimpleHTTPRouter() *SimpleHTTPRouter {
	return &SimpleHTTPRouter{}
}

// WithRoutes inject routes to router
func (router *SimpleHTTPRouter) WithRoutes(routes []Route) {
	router.routes = routes
}

// ServeHTTP handle http request
func (router *SimpleHTTPRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h := router.routeRequest(r)
	if h != nil {
		h.Handle(w, r)
	}
}

func (router *SimpleHTTPRouter) routeRequest(r *http.Request) Handler {
	if router.routes == nil {
		return nil
	}

	var h Handler
	for _, route := range router.routes {
		h = route.MatchHandler(r)
		if h != nil {
			return h
		}
	}

	return nil
}
