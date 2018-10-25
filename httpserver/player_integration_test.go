package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Integration Test
func TestRecordWinAndRetrieveThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := PlayerServer{store}
	player := "Marco"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, getNewScoreRequest(player))

	assertStatus(t, response.Code, http.StatusOK)
	assertBody(t, response.Body.String(), "3")
}
