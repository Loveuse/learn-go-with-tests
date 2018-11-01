package poker

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Integration Test
func TestRecordWinAndRetrieveThem(t *testing.T) {
	database, cleanDatabase := CreateTempFile(t, "[]")
	defer cleanDatabase()

	store, err := NewFileSystemPlayersStore(database)

	if err != nil {
		log.Fatalf("could not connect to the players database %s: %v", dbFileName, err)
	}

	server := NewPlayerServer(store)
	player := "Marco"

	server.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, GetNewScoreRequest(player))

		AssertStatus(t, response.Code, http.StatusOK)
		AssertResponseBody(t, response.Body.String(), "3")
	})
	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, GetLeagueRequest())

		got := DecodeJSONResponse(t, response.Body)
		want := []Player{
			{"Marco", 3},
		}
		AssertLeague(t, got, want)

	})

}
