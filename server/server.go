package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "received wrong method", http.StatusBadRequest)
	}
}

// New generates an instantiated server with necessary handlers in place.
func New() *http.Server {
	s := &http.Server{}

	r := mux.NewRouter()
	r.HandleFunc("/", mainHandler).Methods("GET")

	return s
}

// [ ] listening server (localhost or deployed to Heroku dyno)
// - [ ] listens to render static page with mainHandler
// - - [ ] page contains button to install Slack app
// - [ ] listens for GitHub webhook activity with githubHandler
// - - [ ] if project event received process and send to Slack app
// - [ ] listens for Slack settings update commands with slackHandler
// - - [ ] slackHandler parses incoming slash commands
// [ ] Slack app provides updates to specific channels
// - [ ] provides slash commands to establish settings
