package poker_time

import (
	"log"
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
		nil,
		nil,
	}
	playerServer := NewPlayerServer(&store)
	t.Run("return Marco's score", func(t *testing.T) {

		request := GetNewScoreRequest("Marco")
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusOK)
		AssertResponseBody(t, response.Body.String(), "20")

	})
	t.Run("return Francesco's score", func(t *testing.T) {
		request := GetNewScoreRequest("Francesco")
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusOK)
		AssertResponseBody(t, response.Body.String(), "10")

	})
	t.Run("player missing, 404 status code", func(t *testing.T) {
		request := GetNewScoreRequest("Daniele")
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusNotFound)

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
		request := NewPostWinRequest(player)
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusAccepted)
		AssertPlayerWin(t, &store, player)

	})
}

// Test the LEAGUE function

func TestLeague(t *testing.T) {

	t.Run("it returns the league table as JSON", func(t *testing.T) {

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
		request := GetLeagueRequest()

		playerServer.ServeHTTP(response, request)

		AssertJSONResponse(t, response, jsonContentType)

		var got []Player
		got = DecodeJSONResponse(t, response.Body)

		AssertStatus(t, response.Code, http.StatusOK)
		AssertLeague(t, got, wantedLeague)

	})

}

func TestFileSystemPlayersStore(t *testing.T) {
	database, cleanDatabase := CreateTempFile(t, `[
		{"Name": "Marco", "Score": 20},
		{"Name": "Francesco", "Score": 10}]`)
	defer cleanDatabase()

	store, err := NewFileSystemPlayersStore(database)

	if err != nil {
		log.Fatalf("could not connect to the players database %s: %v", dbFileName, err)
	}

	AssertNoError(t, err)

	t.Run("league from a reader", func(t *testing.T) {

		got := store.GetLeague()

		want := []Player{
			{"Marco", 20},
			{"Francesco", 10},
		}
		AssertLeague(t, got, want)

		got = store.GetLeague()
		AssertLeague(t, got, want)
	})

	t.Run("get player score from a reader", func(t *testing.T) {

		got := store.GetPlayerScore("Marco")
		want := 20

		AssertScore(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		store.RecordWin("Marco")

		got := store.GetPlayerScore("Marco")
		want := 21

		AssertScore(t, got, want)
	})
	t.Run("store wins for new player", func(t *testing.T) {
		store.RecordWin("Daniele")

		got := store.GetPlayerScore("Daniele")
		want := 1

		AssertScore(t, got, want)
	})
	t.Run("work with an empty file", func(t *testing.T) {
		database, cleanDatabase = CreateTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemPlayersStore(database)

		AssertNoError(t, err)
	})
	t.Run("league sorted", func(t *testing.T) {
		database, cleanDatabase = CreateTempFile(t, `[
			{"Name": "Marco", "Score": 20},
			{"Name": "Francesco", "Score": 10}]`)
		defer cleanDatabase()

		store, err = NewFileSystemPlayersStore(database)
		got := store.GetLeague()
		want := []Player{
			{"Marco", 20},
			{"Francesco", 10},
		}

		AssertLeague(t, got, want)
	})
}
