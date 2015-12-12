package model

import (
	"fmt"
)

// Poll is the basic data structure for keeping poll data.
// It has a numeric unique ID, a title/question and a list of possible Answers
type Poll struct {
	ID      uint64
	Title   string   `json:"title"`
	Answers []Answer `json:"answers"`
}

// NewPoll creates a new instance of the struct Poll
func NewPoll(id uint64, title string, answers ...Answer) Poll {
	return Poll{ID: id, Title: title, Answers: answers}
}

func (p Poll) String() string {
	return fmt.Sprintf("Poll #%v, \"%s\"", p.ID, p.Title)
}
