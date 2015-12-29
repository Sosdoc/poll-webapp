package main

import (
	"os"

	"github.com/sosdoc/poll-webapp/controllers"

	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/show", controllers.ReadPoll)
	http.HandleFunc("/create", controllers.WritePoll)
    http.HandleFunc("/vote", controllers.VoteOnPoll)

	http.Handle("/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":"+port, nil)
}
