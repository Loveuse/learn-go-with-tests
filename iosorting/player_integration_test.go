package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Integration Test
func TestRecordWinAndRetrieveThem(t *testing.T) {
	database, cleanDatabase := createTempFile(t, "[]")
	defer cleanDatabase()

	store, err := NewFileSystemPlayersStore(database)

	if err != nil {
		log.Fatalf("could not connect to the players database %s: %v", dbFileName, err)
	}

	server := NewPlayerServer(store)
	player := "Marco"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, getNewScoreRequest(player))

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "3")
	})
	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, getLeagueRequest())

		got := decodeJSONResponse(t, response.Body)
		want := []Player{
			{"Marco", 3},
		}
		assertLeague(t, got, want)

	})

}
