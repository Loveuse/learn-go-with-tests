package poker_time_test

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/loveuse/learn-go-with-tests/time"
)

var dummyBlindAlerter = &SpyBlindAlerter{}
var dummyPlayerStore = &poker_time.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

func TestCLI(t *testing.T) {

	t.Run("start game with 3 players and finish game with 'Marco' as winner", func(t *testing.T) {
		game := &SpyGame{}
		stdout := &bytes.Buffer{}

		in := userSends("3", "Marco wins")
		cli := poker_time.NewCLI(in, stdout, game)

		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, poker_time.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, "Marco")
	})

	t.Run("start game with 8 players and record 'Francesco' as winner", func(t *testing.T) {
		game := &SpyGame{}

		in := userSends("8", "Francesco wins")
		cli := poker_time.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()

		assertGameStartedWith(t, game, 8)
		assertFinishCalledWith(t, game, "Francesco")
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		game := &SpyGame{}

		stdout := &bytes.Buffer{}
		in := userSends("not a number")

		cli := poker_time.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, stdout, poker_time.PlayerPrompt, poker_time.BadPlayerInputErrMsg)
	})
}

func userSends(messages ...string) io.Reader {
	return strings.NewReader(strings.Join(messages, "\n"))
}

// SpyBlindAlerter in order to spy on the blind
type SpyBlindAlerter struct {
	alerts []ScheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, ScheduledAlert{duration, amount})
}

type SpyGame struct {
	StartCalled bool
	StartedWith int

	FinishedCalled bool
	FinishedWith   string
}

func (s *SpyGame) Start(numberOfPlayers int) {
	s.StartCalled = true
	s.StartedWith = numberOfPlayers
}

func (s *SpyGame) Finish(winner string) {
	s.FinishedCalled = true
	s.FinishedWith = winner
}

type ScheduledAlert struct {
	At     time.Duration
	Amount int
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.At)
}

func assertScheduledAlert(t *testing.T, alertGot, alertWant ScheduledAlert) {
	t.Helper()
	amountGot := alertGot.Amount
	if amountGot != alertWant.Amount {
		t.Errorf("amount got %d want %d", amountGot, alertWant.Amount)
	}

	scheduledTimeGot := alertGot.At
	if scheduledTimeGot != alertWant.At {
		t.Errorf("scheduled time got %v want %v", scheduledTimeGot, alertWant.At)
	}
}

func assertMessagesSentToUser(t *testing.T, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got '%s' sent to stdout but expected %+v", got, messages)
	}
}

func assertGameStartedWith(t *testing.T, game *SpyGame, numberOfPlayersWanted int) {
	t.Helper()
	if game.StartedWith != numberOfPlayersWanted {
		t.Errorf("wanted Start called with %d but got %d", numberOfPlayersWanted, game.StartedWith)
	}
}

func assertFinishCalledWith(t *testing.T, game *SpyGame, winner string) {
	t.Helper()
	if game.FinishedWith != winner {
		t.Errorf("expected finish called with '%s' but got '%s'", winner, game.FinishedWith)
	}
}

func assertGameNotStarted(t *testing.T, game *SpyGame) {
	t.Helper()
	if game.StartCalled {
		t.Errorf("game should not have started")
	}
}
