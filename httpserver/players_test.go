package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Testing the GET function
func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Marco":     20,
			"Francesco": 10,
		},
		[]string{},
	}
	playerServer := &PlayerServer{&store}
	t.Run("return Marco's score", func(t *testing.T) {

		request := getNewScoreRequest("Marco")
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertBody(t, response.Body.String(), "20")

	})
	t.Run("return Francesco's score", func(t *testing.T) {
		request := getNewScoreRequest("Francesco")
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertBody(t, response.Body.String(), "10")

	})
	t.Run("player missing", func(t *testing.T) {
		request := getNewScoreRequest("Daniele")
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)

	})
}

// Testing the STORE function
func TestStorePlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		[]string{},
	}
	playerServer := &PlayerServer{&store}

	t.Run("it returns accepted on POST", func(t *testing.T) {
		player := "Marco"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls instead of %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("got winner %s want %s", store.winCalls[0], player)
		}

	})
}

// Stub
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

// Utility functions
func assertBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func getNewScoreRequest(player string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return req
}

func newPostWinRequest(player string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
	return req
}

func assertStatus(t *testing.T, gotStatus, wantStatus int) {
	t.Helper()
	if gotStatus != wantStatus {
		t.Errorf("player should be missing got %d want %d", gotStatus, wantStatus)
	}
}
