package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct {
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

func (i *InMemoryPlayerStore) GetPlayerScore(player string) int {
	return i.store[player]
}

func (i *InMemoryPlayerStore) RecordWin(player string) {
	i.store[player]++
}

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}
	if err := http.ListenAndServe(":4000", server); err != nil {
		log.Fatalf("Could not listen on port 5000 %v", err)
	}
}
