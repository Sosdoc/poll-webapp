package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/valep27/poll-webapp/data"
	"github.com/valep27/poll-webapp/model"
)

// ReadPoll retrieves a Poll and writes it as a JSON object
func ReadPoll(rw http.ResponseWriter, r *http.Request) {
	// TODO: get poll id from request
	query := r.URL.Query()

	hashID := query.Get("id")

	poll, err := data.GetPollByHashID(hashID)
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
	var poll model.Poll
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&poll)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = data.CreatePoll(poll)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("encoding", "text/json")
	rw.Write([]byte("{success: true}"))
}
