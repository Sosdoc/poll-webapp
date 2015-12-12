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

// Answer represents a basic answer to a Poll
// It has an id unique to its Poll and some text associated with it.
// The 'other' flag is used for free form answer, in that case, the text is set by the user.
type Answer struct {
	id    uint32
	Text  string `json:"text"`
	Other bool   `json:"other"`
}

func NewAnswer(id uint32, text string, isOther bool) Answer {
	return Answer{id: id, Text: text, Other: isOther}
}
