package selectp

import (
	"errors"
	"net/http"
	"time"
)

var (
	ErrorTimeout     = errors.New("request timeout")
	tenSecondTimeout = 100 * time.Second
)

func Racer(url1, url2 string) (string, error) {
	return ConfigurableRacer(url1, url2, tenSecondTimeout)
}

func ConfigurableRacer(url1, url2 string, timeout time.Duration) (string, error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout): //1000 * time.Microsecond
		return "", ErrorTimeout
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}
