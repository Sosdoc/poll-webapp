package data

// This package makes use of a PostgreSQL driver, one can be found at https://godoc.org/github.com/lib/pq
// Install it via `go get github.com/lib/pq`

import (
	"database/sql"
	"fmt"
)

// GetDBConnection connects to the database for polls.
// This connections should probably reused when possible.
func GetDBConnection() *sql.DB {
	user := "golang"
	password := "golang"
	dbName := "polls"
	host := "localhost"

	options := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=verify-full",
		user, dbName, password, host)

	db, err := sql.Open("postgres", options)
	if err != nil {
		panic(err)
	}

	return db
}
