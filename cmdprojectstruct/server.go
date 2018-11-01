package poker

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const jsonContentType = "application/json"
const dbFileName = "game.db.json"

type PlayerStore interface {
	GetPlayerScore(player string) int
	RecordWin(player string)
	GetLeague() League
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

type Player struct {
	Name  string
	Score int
}

// constructor for player server: inizialise the server and relative handlers
func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := &PlayerServer{
		store,
		http.NewServeMux(),
	}

	router := http.NewServeMux()
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))

	p.Handler = router
	return p
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	switch r.Method {
	case http.MethodGet:
		p.showScore(w, player)
	case http.MethodPost:
		p.processWin(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(p.store.GetLeague())
	w.Header().Set("content-type", jsonContentType)
	w.WriteHeader(http.StatusOK)
}
