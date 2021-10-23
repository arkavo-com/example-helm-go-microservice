package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// os interrupt
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	// server
	server := http.Server{
		Addr:         "0.0.0.0:80",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	go func() {
		log.Printf("listening on http://%s", server.Addr)
		log.Printf(os.Getenv("SERVICE"))
		log.Printf(os.Getenv("MESSAGE"))
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	<-stop
	err := server.Shutdown(context.Background())
	if err != nil {
		log.Println(err)
	}
}
