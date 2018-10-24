package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://areuserious.mate" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://facebook.com",
		"waat://areuserious.mate",
	}

	want := map[string]bool{
		"http://google.com":       true,
		"http://facebook.com":     true,
		"waat://areuserious.mate": false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func SlowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	numWebsites := 100
	urls := make([]string, numWebsites)

	for i := 0; i < numWebsites; i++ {
		urls[i] = "I am an URL"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(SlowStubWebsiteChecker, urls)
	}
}
