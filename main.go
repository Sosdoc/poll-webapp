package main

import (
	"encoding/json"
	"github.com/sosdoc/polls/model"
	"os"

	"net/http"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/poll", writePoll)
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":"+port, nil)
}

func writePoll(rw http.ResponseWriter, r *http.Request) {
	yes := model.NewAnswer(uint32(0), "Yes", false)
	no := model.NewAnswer(uint32(1), "No", false)
	poll := model.NewPoll("Do you like turtles?", yes, no)

	js, err := json.Marshal(poll)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("encoding", "text/json")
	rw.Write(js)
}
