package model

import (
	"fmt"
)

// Poll is the basic data structure for keeping poll data.
// It has a numeric unique ID, a title/question and a list of possible Answers
type Poll struct {
	id      uint64
	title   string
	answers []Answer
}

func NewPoll(title string, answers ...Answer) Poll {
	id := uint64(0)
	return Poll{id: id, title: title, answers: answers}
}

func (p Poll) String() string {
	return fmt.Sprintf("Poll #%v, \"%s\"", p.id, p.title)
}

// Answer represents a basic answer to a Poll
// It has an id unique to its Poll and some text associated with it.
// The 'other' flag is used for free form answer, in that case, the text is set by the user.
type Answer struct {
	id    uint32
	text  string
	other bool
}

func NewAnswer(id uint32, text string, isOther bool) Answer {
	return Answer{id: id, text: text, other: isOther}
}
