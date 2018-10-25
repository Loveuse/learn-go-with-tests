package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

const jsonContentType = "application/json"

// Testing the GET function
func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Marco":     20,
			"Francesco": 10,
		},
		[]string{},
		[]Player{},
	}
	playerServer := NewPlayerServer(&store)
	t.Run("return Marco's score", func(t *testing.T) {

		request := getNewScoreRequest("Marco")
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")

	})
	t.Run("return Francesco's score", func(t *testing.T) {
		request := getNewScoreRequest("Francesco")
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")

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
		[]Player{},
	}
	playerServer := NewPlayerServer(&store)

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

// Test the LEAGUE function

func TestLeague(t *testing.T) {

	t.Run("it returns 200 on /league", func(t *testing.T) {

		wantedLeague := []Player{
			{"Marco", 20},
			{"Francesco", 10},
		}

		store := StubPlayerStore{
			nil,
			nil,
			wantedLeague,
		}
		playerServer := NewPlayerServer(&store)

		response := httptest.NewRecorder()
		request := getLeagueRequest()

		playerServer.ServeHTTP(response, request)

		assertJSONResponse(t, response, jsonContentType)

		var got []Player
		got = decodeJSONResponse(t, response.Body)

		assertStatus(t, response.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)

	})

}

// Stub
type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (s *StubPlayerStore) GetPlayerScore(player string) int {
	score := s.scores[player]
	return score
}

func (s *StubPlayerStore) RecordWin(player string) {
	s.winCalls = append(s.winCalls, player)
}

func (s *StubPlayerStore) GetLeague() []Player {
	return s.league
}

// Utility functions
func assertResponseBody(t *testing.T, got, want string) {
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

func assertLeague(t *testing.T, gotLeague, wantedLeague []Player) {
	t.Helper()
	if !reflect.DeepEqual(gotLeague, wantedLeague) {
		t.Errorf("got league %v wanted %v", gotLeague, wantedLeague)
	}
}

func assertJSONResponse(t *testing.T, response *httptest.ResponseRecorder, want string) {
	if response.Header().Get("content-type") != want {
		t.Errorf("response did not have content-type of application/json, got %v", response.HeaderMap)
	}
}

func getLeagueRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
}

func decodeJSONResponse(t *testing.T, body io.Reader) (league []Player) {
	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Fatalf("Unable to parse response '%s' into slice of Player, %v", body, err)
	}

	return league
}
