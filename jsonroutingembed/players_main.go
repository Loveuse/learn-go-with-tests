package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer(NewInMemoryPlayerStore())
	if err := http.ListenAndServe(":4000", server); err != nil {
		log.Fatalf("Could not listen on port 5000 %v", err)
	}
}
