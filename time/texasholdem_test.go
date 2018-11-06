package poker_time_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/loveuse/learn-go-with-tests/time"
)

func TestGame_Start(t *testing.T) {
	t.Run("schedules alerts on game start for 5 players", func(t *testing.T) {
		blindAlerter := &SpyBlindAlerter{}
		game := poker_time.NewTexasHoldem(dummyPlayerStore, blindAlerter)

		game.Start(5)

		cases := []ScheduledAlert{
			{At: 0 * time.Second, Amount: 100},
			{At: 10 * time.Minute, Amount: 200},
			{At: 20 * time.Minute, Amount: 300},
			{At: 30 * time.Minute, Amount: 400},
			{At: 40 * time.Minute, Amount: 500},
			{At: 50 * time.Minute, Amount: 600},
			{At: 60 * time.Minute, Amount: 800},
			{At: 70 * time.Minute, Amount: 1000},
			{At: 80 * time.Minute, Amount: 2000},
			{At: 90 * time.Minute, Amount: 4000},
			{At: 100 * time.Minute, Amount: 8000},
		}

		checkSchedulingCases(t, cases, blindAlerter)
	})

	t.Run("schedules alerts on game start for 7 players", func(t *testing.T) {
		blindAlerter := &SpyBlindAlerter{}
		game := poker_time.NewTexasHoldem(dummyPlayerStore, blindAlerter)

		game.Start(7)

		cases := []ScheduledAlert{
			{At: 0 * time.Second, Amount: 100},
			{At: 12 * time.Minute, Amount: 200},
			{At: 24 * time.Minute, Amount: 300},
			{At: 36 * time.Minute, Amount: 400},
		}

		checkSchedulingCases(t, cases, blindAlerter)
	})

}

func TestGame_Finish(t *testing.T) {
	store := &poker_time.StubPlayerStore{}
	game := poker_time.NewTexasHoldem(store, dummyBlindAlerter)
	winner := "Ruth"

	game.Finish(winner)
	poker_time.AssertPlayerWin(t, store, winner)
}

func checkSchedulingCases(t *testing.T, cases []ScheduledAlert, blindAlerter *SpyBlindAlerter) {
	t.Helper()
	var alertGot ScheduledAlert
	for i, c := range cases {
		t.Run(fmt.Sprintf("%d scheduled for %v", c.Amount, c.At), func(t *testing.T) {
			if len(blindAlerter.alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
			}

			alertGot = blindAlerter.alerts[i]
			assertScheduledAlert(t, alertGot, c)
		})
	}

}
