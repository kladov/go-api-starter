package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// IndexRoute
type IndexRoute struct{}

// MatchHandler
func (h *IndexRoute) MatchHandler(r *http.Request) http.HandlerFunc {
	if r.URL.Path == "/" {
		switch r.Method {
		case http.MethodGet:
			return IndexGetHandler
		case http.MethodPost:
			return IndexPostHandler
		case http.MethodPut:
			return IndexPutHandler
		case http.MethodDelete:
			return IndexDeleteHandler
		}

		return NotFoundHandler
	}

	return nil
}

// IndexGetHandler api root for get requests
var IndexGetHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	v := r.URL.Query().Get("field")
	json.NewEncoder(w).Encode(Response{"Hello world (get) with param value: '" + v + "'"})
}

type indexRequestData struct {
	Field string `json:"field"`
}

// IndexPostHandler api root for post requests
var IndexPostHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	rData, err := extractData(r)
	if err != nil {
		json.NewEncoder(w).Encode(Response{"Hello world (post) error: " + err.Error()})
		return
	}
	json.NewEncoder(w).Encode(Response{"Hello world (post) param value: " + rData.Field})
}

// IndexPutHandler api root for put requests
var IndexPutHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	rData, err := extractData(r)
	if err != nil {
		json.NewEncoder(w).Encode(Response{"Hello world (put) error: " + err.Error()})
		return
	}
	json.NewEncoder(w).Encode(Response{"Hello world (put) param value: " + rData.Field})
}

// IndexDeleteHandler api root delete requests
var IndexDeleteHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	rData, err := extractData(r)
	if err != nil {
		json.NewEncoder(w).Encode(Response{"Hello world (delete) error: " + err.Error()})
		return
	}
	json.NewEncoder(w).Encode(Response{"Hello world (delete) param value: " + rData.Field})
}

func extractData(r *http.Request) (*indexRequestData, error) {
	rData := &indexRequestData{}
	switch r.Header.Get("Content-type") {
	case "application/json":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(body, rData)
	case "application/x-www-form-urlencoded":
		r.ParseForm()
		rData.Field = r.Form.Get("field")
	default:
		r.ParseMultipartForm(0)
		rData.Field = r.FormValue("field")
	}

	return rData, nil
}
