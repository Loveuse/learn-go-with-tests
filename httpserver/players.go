package main

import (
	"fmt"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(player string) int
	RecordWin(player string)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	switch r.Method {
	case http.MethodGet:
		p.showScore(w, player)
	case http.MethodPost:
		p.processWin(w, player)
	}
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(player string) int {
	score := s.scores[player]
	return score
}

func (s *StubPlayerStore) RecordWin(player string) {
	s.winCalls = append(s.winCalls, player)
}

func GetPlayerScore(player string) int {

	if player == "Marco" {
		return 20
	}

	if player == "Francesco" {
		return 10
	}

	return 0

}
