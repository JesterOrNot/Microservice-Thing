package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	handlers "product-api/handlers"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "product-api: ", log.LstdFlags)

	hh := handlers.NewHello(logger)
	gh := handlers.NewGoodbye(logger)

	sm := http.NewServeMux()
	sm.Handle("/hello", hh)
	sm.Handle("/goodbye", gh)

	// Timeouts are used to prevent DOS (denial of service) attacks
	s := &http.Server{
		Addr:         ":8000",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	fmt.Println("Server listening on Port 8000")
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt)
	signal.Notify(channel, os.Kill)
	sig := <-channel
	logger.Println("Recieved Graceful shutdown", sig)

	// Allow graceful shutdown
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
