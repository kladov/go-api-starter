package mux

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSimpleHTTPRouter_ServeHTTPMiddlewareOrder(t *testing.T) {
	r := NewSimpleHTTPRouter()
	r.WithRoutes([]Route{NewRouteStub()})
	r.WithMiddleware([]Middleware{mw1, mw2})

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if rr.Body.String() != "mw1->mw2->h1" {
		t.Errorf("Middleware call ordering is wrong. It should be 'mw1->mw2->h1', '%s' given", rr.Body.String())
	}
}

func TestSimpleHTTPRouter_ServeHTTPMiddlewareBrake(t *testing.T) {
	r := NewSimpleHTTPRouter()
	r.WithRoutes([]Route{NewRouteStub()})
	r.WithMiddleware([]Middleware{mw1, mwBrake})

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if rr.Body.String() != "mw1->mwBrake " {
		t.Errorf("Middleware brake not wirk. It should be 'mw1->mwBrake ', '%s' given", rr.Body.String())
	}
}

var mw1 Middleware = func(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("mw1->"))
		next.ServeHTTP(w, r)
	}
}

var mw2 Middleware = func(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("mw2->"))
		next.ServeHTTP(w, r)
	}
}

var mwBrake Middleware = func(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("mwBrake "))
	}
}

var h1 http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("h1"))
}

type RouteMock struct {
	MatchFunc func(r *http.Request) http.HandlerFunc
}

func NewRouteMock(matchFunc func(r *http.Request) http.HandlerFunc) Route {
	return &RouteMock{matchFunc}
}

func NewRouteStub() Route {
	return NewRouteMock(func(r *http.Request) http.HandlerFunc {
		return h1
	})
}

func (routeMock *RouteMock) MatchHandler(r *http.Request) http.HandlerFunc {
	return routeMock.MatchFunc(r)
}
