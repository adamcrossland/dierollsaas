// derolsaas.go
package main

import (
	"gorilla/mux"
	"net/http"
	"roller"
)

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/roll/{roll}", RollHandler)
	http.Handle("/", r)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//mtemplate.RenderFile("html/index.html", w, nil)
}

func RollHandler(w http.ResponseWriter, r *http.Request) {
	//c := NewContext(r)
	vars := mux.Vars(r)
	rollRequest := vars["roll"]
	if len(rollRequest) == 0 {
		// handle error message
	}
	spec, specErr := roller.Parse(rollRequest)
	if specErr != nil {
		// handle error message
	}
	results := roller.DoRolls(*spec)
	w.Write(results.ToJSON())
}
