package controllers

import (
	"encoding/json"
	"github.com/sosdoc/polls/data"
	"github.com/sosdoc/polls/model"
	"net/http"
)

// ReadPoll retrieves a Poll and writes it as a JSON object
func ReadPoll(rw http.ResponseWriter, r *http.Request) {
	// TODO: get poll id from request
	pollID := uint64(0)
	poll, err := data.GetPollByID(pollID)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(poll)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("encoding", "text/json")
	rw.Write(js)
}

// WritePoll receives a JSON object representing a new Poll.
// It will try to insert the poll in the data storage.
func WritePoll(rw http.ResponseWriter, r *http.Request) {
	// TODO: get the poll from form data
	yes := model.NewAnswer(uint32(0), "Yes", false)
	no := model.NewAnswer(uint32(1), "No", false)
	poll := model.NewPoll(uint64(0), "Do you like turtles?", yes, no)

	err := data.CreatePoll(poll)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("encoding", "text/json")
	rw.Write([]byte("{success: true}"))
}
