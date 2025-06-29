package main

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)

	http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log.SetFormatter(&log.TextFormatter{
			FullTimestamp: true,
		})
		log.SetLevel(log.DebugLevel)

		select {
		case <-ctx.Done():
			log.Debugf("Context got cancelled")
		case <-time.After(3 * time.Second):
			log.Info("Processed the request in 3 seconds")
			fmt.Fprintln(w, "200 OK response")
		}
	})
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Warn("Can't start the server %v", err.Error())
	}
}
