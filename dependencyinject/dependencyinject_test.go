package dependencyinject

import (
	"bytes"
	"testing"
)

func TestGreetings(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Marco")

	got := buffer.String()
	want := "Hello, Marco"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
