package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/yashasn/game-api/server"
)

type InMemoryStore struct {
}

func (s *InMemoryStore) GetPlayerScore(name string) int {
	return 1
}

func main() {
	/*The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers.
	If f is a function with the appropriate signature, HandlerFunc(f) is a Handler that calls f */
	//handler := http.HandlerFunc(PlayerServer)
	store := InMemoryStore{}
	server := &server.PlayerServer{Store: &store}
	//This will start a web server listening on a port, creating a goroutine for every request and running it against a handler
	port, err := getPort()
	if err != nil {
		fmt.Println("Error parsing PORT environment variable:", err)
		os.Exit(1)
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), server))
}

func getPort() (int, error) {
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080" // Default port
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return 0, err
	}
	return port, nil
}
