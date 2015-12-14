package data

// This package makes use of a PostgreSQL driver, one can be found at https://godoc.org/github.com/lib/pq
// Install it via `go get github.com/lib/pq`

import (
	"database/sql"
	"fmt"

	"github.com/sosdoc/poll-webapp/model"
)

var dbInstance *sql.DB

// getDBConnection connects to the database for polls.
// If the connection was created already, it will return the previous reference.
func getDBConnection() *sql.DB {
	if dbInstance != nil {
		return dbInstance
	}

	user := "golang"
	password := "golang"
	dbName := "polldb"
	host := "localhost"

	options := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable",
		user, dbName, password, host)

	db, err := sql.Open("postgres", options)
	if err != nil {
		panic(err)
	}

	dbInstance = db

	return db
}

// CreatePoll inserts the specified poll in the data storage.
func CreatePoll(p model.Poll) error {
	db := getDBConnection()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// pollID is needed to connect answers to the poll
	var pollID int

	err = tx.QueryRow("INSERT INTO polls (poll_id, title) VALUES (DEFAULT, $1) RETURNING poll_id",
		p.Title).Scan(&pollID)
	if err != nil {
		return err
	}

	for _, a := range p.Answers {
		_, err = tx.Exec("INSERT INTO answer VALUES ($1, $2, $3, $4)",
			a.ID, pollID, a.Text, false)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// GetPollByHashID retrieves a Poll from the data storage using its unique hash ID.
func GetPollByHashID(hash string) (p *model.Poll, err error) {
	db := getDBConnection()
	pollID := model.DecodePollID(hash)

	var pollTitle string
	err = db.QueryRow("SELECT title FROM polls WHERE poll_id=$1", pollID).Scan(&pollTitle)

	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT id, title, other FROM answer WHERE poll_id=$1", pollID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var answers []model.Answer

	for rows.Next() {
		var (
			answerID int
			title    string
			other    bool
		)

		err = rows.Scan(&answerID, &title, &other)
		answers = append(answers, model.NewAnswer(answerID, title, other))
	}

	poll := model.NewPollWithHash(hash, pollTitle, answers...)

	return &poll, nil
}
