package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	handlers "product-api/handlers"
)

func main() {
	logger := log.New(os.Stdout, "product-api: ", log.LstdFlags)
	
	hh := handlers.NewHello(logger)
	gh := handlers.NewGoodbye(logger)

	sm := http.NewServeMux()
	sm.Handle("/hello", hh)
	sm.Handle("/goodbye", gh)

	fmt.Println("Server listening on Port 8000")
	http.ListenAndServe(":8000", sm)
}
