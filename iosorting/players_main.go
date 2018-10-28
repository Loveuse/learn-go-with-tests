package main

import (
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s: %v", dbFileName, err)
	}

	store, err := NewFileSystemPlayersStore(db)

	if err != nil {
		log.Fatalf("could not connect to the players database %s: %v", dbFileName, err)
	}

	server := NewPlayerServer(store)

	if err := http.ListenAndServe(":4000", server); err != nil {
		log.Fatalf("Could not listen on port 5000 %v", err)
	}
}
