package model

import (
	"fmt"
)

// Poll is the basic data structure for keeping poll data.
// It has a numeric unique ID, a title/question and a list of possible Answers
type Poll struct {
	id      uint64
	Title   string   `json:"title"`
	Answers []Answer `json:"answers"`
}

func NewPoll(title string, answers ...Answer) Poll {
	id := uint64(0)
	return Poll{id: id, Title: title, Answers: answers}
}

func (p Poll) String() string {
	return fmt.Sprintf("Poll #%v, \"%s\"", p.id, p.Title)
}
