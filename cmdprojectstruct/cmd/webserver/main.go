package main

import (
	"log"
	"net/http"

	"github.com/loveuse/learn-go-with-tests/cmdprojectstruct"
)

const dbFileName = "game.db.json"

func main() {

	store, err := poker.FileSystemPlayersStoreFromFile(dbFileName)
	if err != nil {
		log.Fatalf("could not set up the store: %v", err)
	}

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":4000", server); err != nil {
		log.Fatalf("Could not listen on port 5000 %v", err)
	}
}
