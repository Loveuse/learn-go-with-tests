package poker_time

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

// Stub
type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   League
}

func (s *StubPlayerStore) GetPlayerScore(player string) int {
	score := s.scores[player]
	return score
}

func (s *StubPlayerStore) RecordWin(player string) {
	s.winCalls = append(s.winCalls, player)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

// Utility functions
func AssertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func GetNewScoreRequest(player string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return req
}

func NewPostWinRequest(player string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
	return req
}

func AssertStatus(t *testing.T, gotStatus, wantStatus int) {
	t.Helper()
	if gotStatus != wantStatus {
		t.Errorf("player should be missing got %d want %d", gotStatus, wantStatus)
	}
}

func AssertLeague(t *testing.T, gotLeague, wantedLeague []Player) {
	t.Helper()
	if !reflect.DeepEqual(gotLeague, wantedLeague) {
		t.Errorf("got league %v wanted %v", gotLeague, wantedLeague)
	}
}

func AssertJSONResponse(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Header().Get("content-type") != want {
		t.Errorf("response did not have content-type of application/json, got %v", response.HeaderMap)
	}
}

func GetLeagueRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
}

func DecodeJSONResponse(t *testing.T, body io.Reader) (league []Player) {
	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Fatalf("unable to parse response '%s' into slice of Player, %v", body, err)
	}

	return league
}

func AssertScore(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got score %d want %d", got, want)
	}
}

func CreateTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()

	tmpFile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create a temp file %v", err)
	}

	tmpFile.Write([]byte(initialData))
	removeFile := func() {
		os.Remove(tmpFile.Name())
	}

	return tmpFile, removeFile
}

func AssertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didnt expect an error but got one, %v", err)
	}
}

func AssertPlayerWin(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()
	if len(store.winCalls) > 1 {
		t.Errorf("got %d calls instead of %d", len(store.winCalls), 1)
	}

	if store.winCalls[0] != winner {
		t.Errorf("got winner %s want %s", store.winCalls[0], winner)
	}
}

func FileSystemPlayersStoreFromFile(path string) (*FileSystemPlayersStore, error) {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		return nil, fmt.Errorf("problem opening %s: %v", dbFileName, err)
	}

	store, err := NewFileSystemPlayersStore(db)

	if err != nil {
		return nil, fmt.Errorf("could not connect to the players database %s: %v", dbFileName, err)
	}

	return store, nil
}
