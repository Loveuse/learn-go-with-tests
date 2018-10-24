package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	countdownStart = 3
	finalWorld     = "Go!"
)

type Sleeper interface {
	Sleep()
}

// type SpySleeper struct {
// 	Calls int
// }

// func (s *SpySleeper) Sleep() {
// 	s.Calls++
// }

// type DefaultSleeper struct{}

// func (d *DefaultSleeper) Sleep() {
// 	time.Sleep(1 * time.Second)
// }

const (
	write = "write"
	sleep = "sleep"
)

type CountdownOperationSpy struct {
	Calls []string
}

func (c *CountdownOperationSpy) Sleep() {
	c.Calls = append(c.Calls, sleep)
}

func (c *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, write)
	return
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	DurationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.DurationSlept = duration
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(writer, i)
	}
	sleeper.Sleep()
	fmt.Fprint(writer, finalWorld)
}

func main() {
	//sleeper := &CountdownOperationSpy{}
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
