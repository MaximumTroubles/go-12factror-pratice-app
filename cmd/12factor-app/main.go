package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/gorilla/mux"
)

func main() {
	log.Info("Hello world")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Fatal("Port is not set")
	}

	// DB_URL heroku key if we want to use database.

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	serv := http.Server{
		Addr: net.JoinHostPort("", port),
		Handler: router,
	}

	go serv.ListenAndServe()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	<-interrupt

	timeout, cancelFunc := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancelFunc()

	serv.Shutdown(timeout)	
}
