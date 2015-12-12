package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/sosdoc/polls/model"
)

func ReadPoll(rw http.ResponseWriter, r *http.Request) {
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

func WritePoll(rw http.ResponseWriter, r *http.Request) {

}
