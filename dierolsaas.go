// dierolsaas.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"roller"

	"bitbucket.org/adamcrossland/mtemplate"
	"github.com/gorilla/mux"
)

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/index", indexHandler)
	r.HandleFunc("/roll/{roll}", rollHandler)
	r.HandleFunc("/slack/roll/{roll}", slackHandler)
	http.Handle("/", r)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	mtemplate.RenderFile("html/index.html", w, nil)
}

func doRollRequest(w http.ResponseWriter, r *http.Request) (*roller.RollResults, bool) {
	vars := mux.Vars(r)
	rollRequest := vars["roll"]
	if len(rollRequest) == 0 {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Missing roll specification")
		return nil, false
	}

	spec, specErr := roller.Parse(rollRequest)
	if specErr != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Roll specification in incorrect format: %s", specErr)
		return nil, false
	}

	rollResult := roller.DoRolls(*spec)
	return &rollResult, true
}

func rollHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	results, ok := doRollRequest(w, r)

	switch r.URL.RawQuery {
	case "text":
		if ok {
			w.Header().Set("Content-Type", "text/plain")
			w.Write(results.ToText())
		}
	case "json":
		fallthrough
	default:
		w.Header().Set("Content-Type", "application/json")
		if ok {
			w.Write(results.ToJSON())
		}
	}
}

const slackToken string = "IbAZE0ckwoSNJcsGWE7sqX5j"

type slackResponse struct {
	Text         string `json:"text"`
	ResponseType string `json:"response_type"`
}

func slackHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	foundToken := r.PostFormValue("token")

	// Make sure that this is a legit request from Slack
	if foundToken == slackToken {
		// Make sure that we have a repsonse url to which to write
		respondTo := r.PostFormValue("response_url")
		if respondTo != "" {
			results, ok := doRollRequest(w, r)
			var slack slackResponse
			if ok {
				slack.Text = results.String()
				slack.ResponseType = "in_channel"

				coded, _ := json.Marshal(slack)
				w.Write(coded)
			}
		} else {
			w.WriteHeader(400)
			fmt.Fprintf(w, "Missing responseurl")
		}
	} else if foundToken == "" {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Missing token")
	} else {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Invalid token")
	}
}
