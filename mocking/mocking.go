package mocking

import (
	"fmt"
	"io"
	"time"
)

const (
	countdownStart = 3
	finalWord      = "Go"
	sleep          = "sleep"
	write          = "write"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpySleeperOperations struct {
	Calls []string
}

func (s *SpySleeperOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpySleeperOperations)

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	fmt.Fprintf(writer, finalWord)
}
