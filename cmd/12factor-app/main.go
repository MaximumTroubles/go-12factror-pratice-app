package main

import (
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

}
