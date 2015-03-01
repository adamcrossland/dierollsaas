// derolsaas.go
package main

import (
	"gorilla/mux"
	"mtemplate"
	"net/http"
)

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/roll/{roll}", RollHandler)
	http.Handle("/", r)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	mtemplate.RenderFile("html/index.html", w, nil)
}

func RollHandler(w http.ResponseWriter, r *http.Request) {
	c := NewContext(r)
}
