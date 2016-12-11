// derolsaas.go
package main

import (
	"bytes"
	"fmt"
	"net/http"

	"roller"

	"bitbucket.org/adamcrossland/mtemplate"
	"github.com/gorilla/mux"
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
	r.ParseForm()

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
		fmt.Fprintf(w, "Roll specification in incorrect format: %s", specErr)
		return
	}
	results := roller.DoRolls(*spec)

	switch r.URL.RawQuery {
	case "slack":
	case "text":
		var buf bytes.Buffer
		for er := 0; er < results.Count; er++ {
			if er > 0 {
				buf.WriteString("; ")
			}
			buf.WriteString(fmt.Sprintf("%d", results.Rolls[er].Total))
			if results.Rolls[er].Count > 1 {
				for ed := 0; ed < results.Rolls[er].Count; ed++ {
					if ed == 0 {
						buf.WriteString(": ")
					} else {
						buf.WriteString(" ")
					}
					buf.WriteString(fmt.Sprintf("[%d]", results.Rolls[er].Dies[ed]))
				}
			}
		}
		w.Header().Set("Content-Type", "text/plain")
		w.Write(buf.Bytes())
	default:
		w.Header().Set("Content-Type", "application/json")
		w.Write(results.ToJSON())
	}
}
