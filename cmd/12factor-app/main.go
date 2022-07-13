package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Hello world")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Fatal("Port is not set")
	}

	// DB_URL heroku key if we want to use database.

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	http.ListenAndServe(":"+port, nil)
}
