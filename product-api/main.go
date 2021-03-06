package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	handlers "product-api/handlers"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "product-api: ", log.LstdFlags)

	ph := handlers.NewProducts(logger)

	sm := http.NewServeMux()
	sm.Handle("/api/products", ph)

	// Timeouts are used to prevent DOS (denial of service) attacks
	s := &http.Server{
		Addr:         ":8000",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	logger.Println("Server listening on Port 8000")
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
	logger.Println("Received Graceful shutdown", sig)

	// Allow graceful shutdown
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
