package model

// Answer represents a basic answer to a Poll
// It has an id unique to its Poll and some text associated with it.
// The 'other' flag is used for free form answer, in that case, the text is set by the user.
type Answer struct {
	ID    uint32
	Text  string `json:"text"`
	Other bool   `json:"other"`
}

// NewAnswer creates a new instance of the struct Answer.
func NewAnswer(id uint32, text string, isOther bool) Answer {
	return Answer{ID: id, Text: text, Other: isOther}
}
