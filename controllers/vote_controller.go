package controllers

import (
	"net/http"
	"strconv"

	"github.com/valep27/poll-webapp/data"
)

// VoteOnPoll retrieves a Poll and writes it as a JSON object
func VoteOnPoll(rw http.ResponseWriter, r *http.Request) {
	// TODO: get poll id from request
	query := r.URL.Query()

	hashID := query.Get("poll_id")
	answerID, err := strconv.Atoi(query.Get("answer_id"))

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = data.InsertVote(hashID, answerID)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("encoding", "text/json")
	rw.Write([]byte("{success: true}"))
}
