package main

import (
	"github.com/sosdoc/polls/controllers"
	"os"

	// postgres driver
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/poll", controllers.ReadPoll)
	http.HandleFunc("/write", controllers.WritePoll)

	http.Handle("/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":"+port, nil)
}
