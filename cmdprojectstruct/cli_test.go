package poker_test

import (
	"strings"
	"testing"

	"github.com/loveuse/learn-go-with-tests/cmdprojectstruct"
)

func TestCLI(t *testing.T) {

	t.Run("record Marco victory", func(t *testing.T) {
		in := strings.NewReader("Marco wins\n")
		store := &poker.StubPlayerStore{}

		cli := poker.NewCLI(store, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, store, "Marco")
	})
	t.Run("record Francesco victory", func(t *testing.T) {
		in := strings.NewReader("Francesco wins\n")
		store := &poker.StubPlayerStore{}

		cli := poker.NewCLI(store, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, store, "Francesco")
	})

}
