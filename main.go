package main

import (
	"github.com/sosdoc/polls/controllers"
	"os"

	"net/http"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/poll", controllers.ReadPoll)

	http.Handle("/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":"+port, nil)
}
