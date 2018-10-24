package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("prints 3 2 1 Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &CountdownOperationSpy{}

		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("sleep after every print", func(t *testing.T) {
		countdownOperationSpy := &CountdownOperationSpy{}
		Countdown(countdownOperationSpy, countdownOperationSpy)

		got := countdownOperationSpy.Calls
		want := []string{sleep, write, sleep, write, sleep, write, sleep, write}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("different order of sleeps got %s want %s", got, want)
		}

	})
}

func TestConfigurableSleep(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	configurableSleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	configurableSleeper.Sleep()

	if spyTime.DurationSlept != sleepTime {
		t.Errorf("Incorrect time slept got %v want %v", spyTime.DurationSlept, sleepTime)
	}

}
