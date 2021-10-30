package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"example-helm-go-microservice/pkg/pet"
	"github.com/jackc/pgx/v4"
)

func main() {
	// Open up our database connection.
	config, err := pgx.ParseConfig("postgres://host:5432/database?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	config.Host = os.Getenv("POSTGRES_HOST")
	config.Database = os.Getenv("POSTGRES_DATABASE")
	config.User = os.Getenv("POSTGRES_USER")
	config.Password = "mysecretpassword"
	config.LogLevel = pgx.LogLevelTrace
	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}
	//defer the close till after the main function has finished	executing
	defer conn.Close(context.Background())
	var greeting string
	//
	conn.QueryRow(context.Background(), "select 1").Scan(&greeting)
	fmt.Println(greeting)

	// os interrupt
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	// server
	server := http.Server{
		Addr:         "0.0.0.0:8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	http.HandleFunc("/", pet.Handler)
	go func() {
		log.Printf("listening on http://%s", server.Addr)
		log.Printf(os.Getenv("SERVICE"))
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	<-stop
	err = server.Shutdown(context.Background())
	if err != nil {
		log.Println(err)
	}
}
