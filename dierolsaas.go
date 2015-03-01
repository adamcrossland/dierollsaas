// derolsaas.go
package main

import (
	"fmt"
	"gorilla/mux"
	"mtemplate"
	"net/http"
	"roller"
)

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/index", IndexHandler)
	r.HandleFunc("/roll/{roll}", RollHandler)
	http.Handle("/", r)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	mtemplate.RenderFile("html/index.html", w, nil)
}

func RollHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rollRequest := vars["roll"]
	if len(rollRequest) == 0 {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Missing roll specification")
		return
	}
	spec, specErr := roller.Parse(rollRequest)
	if specErr != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Roll specification in incorrect format")
		return
	}
	results := roller.DoRolls(*spec)
	w.Header().Set("Content-Type", "application/json")
	w.Write(results.ToJSON())
}
