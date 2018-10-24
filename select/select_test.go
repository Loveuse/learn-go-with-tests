package selectp

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("let's win", func(t *testing.T) {
		slowServer := makeDelayedServer(1 * time.Millisecond)
		fastServer := makeDelayedServer(0)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		got, _ := Racer(slowURL, fastURL)
		want := fastURL

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
	t.Run("let's have an error for timeout", func(t *testing.T) {
		timeoutServer := makeDelayedServer(10 * time.Millisecond)

		defer timeoutServer.Close()

		_, err := ConfigurableRacer(timeoutServer.URL, timeoutServer.URL, 5*time.Millisecond)

		if err == nil {
			t.Errorf("got %s want %s", err, ErrorTimeout)
		}
	})

}

func makeDelayedServer(sleepTime time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(sleepTime)
		w.WriteHeader(http.StatusOK)
	}))
}
