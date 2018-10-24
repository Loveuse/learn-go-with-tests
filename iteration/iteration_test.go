package iteration

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {

	t.Run("say hello to world", func(t *testing.T) {
		got := Hello()
		want := "Hello, World!"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

}

func ExampleReverse() {
	fmt.Println("olleH")
	// Output: olleH
}

func TestIteration(t *testing.T) {
	got := Repeat("a")
	want := "aaaaa"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
