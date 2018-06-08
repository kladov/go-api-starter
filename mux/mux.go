package mux

import (
	"net/http"
)

// Route go-between Handler and http.Request
type Route interface {
	// MatchHandler return related Handler based on http.Request (nil if not matched)
	MatchHandler(r *http.Request) http.HandlerFunc
}

// Middleware run before request handled
type Middleware func(next http.HandlerFunc) http.HandlerFunc

// SimpleHTTPRouter simple implementation of HTTP request router
type SimpleHTTPRouter struct {
	routes     []Route
	middleware []Middleware
}

// NewSimpleHTTPRouter crete new instance of SimpleHTTPRouter
func NewSimpleHTTPRouter() *SimpleHTTPRouter {
	return &SimpleHTTPRouter{}
}

// WithRoutes inject routes to router
func (router *SimpleHTTPRouter) WithRoutes(routes []Route) {
	router.routes = routes
}

// WithMiddleware inject middleware to router
func (router *SimpleHTTPRouter) WithMiddleware(m []Middleware) {
	router.middleware = m
}

// ServeHTTP handle http request
func (router *SimpleHTTPRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h := router.routeRequest(r)

	if h == nil {
		return
	}

	if len(router.middleware) > 0 {
		for i := len(router.middleware) - 1; i >= 0; i-- {
			h = router.middleware[i](h)
		}
	}

	h.ServeHTTP(w, r)
}

func (router *SimpleHTTPRouter) routeRequest(r *http.Request) http.HandlerFunc {
	if router.routes == nil {
		return nil
	}

	var h http.HandlerFunc
	for _, route := range router.routes {
		h = route.MatchHandler(r)
		if h != nil {
			return h
		}
	}

	return nil
}
