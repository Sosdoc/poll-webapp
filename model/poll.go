package model

import (
	"fmt"

	"github.com/speps/go-hashids"
)

// Poll is the basic data structure for keeping poll data.
// It has a unique ID string, a title/question and a list of possible Answers
type Poll struct {
	ID      string
	Title   string   `json:"title"`
	Answers []Answer `json:"answers"`
}

// NewPoll creates a new instance of the struct Poll
func NewPoll(title string, answers ...Answer) Poll {
	return Poll{Title: title, Answers: answers}
}

// NewPollWithHash creates a new instance of the struct Poll with the provided hash ID
func NewPollWithHash(hash string, title string, answers ...Answer) Poll {
	p := NewPoll(title, answers...)
	p.ID = hash
	return p
}

// EncodePollID transforms an integer poll_id compatible with the data layer into
// an hashID, a.k.a. a unique short string that can be used on the frontend and
// decoded back into the original integer.
func EncodePollID(id int) (string, error) {
	h := hashids.NewWithData(getHashIDData())
	return h.Encode([]int{id})
}

// DecodePollID transforms an hashID to a unique poll integer id.
func DecodePollID(hash string) int {
	h := hashids.NewWithData(getHashIDData())
	return h.Decode(hash)[0]
}

// getHashIDData creates the data used to encode/decode poll ids.
// The salt and lenght are also defined here.
func getHashIDData() *hashids.HashIDData {
	hd := hashids.NewData()
	hd.Salt = "poll-webapp"
	hd.MinLength = 8
	return hd
}

func (p Poll) String() string {
	return fmt.Sprintf("Poll #%v, \"%s\"", p.ID, p.Title)
}
