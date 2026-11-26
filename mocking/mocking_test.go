package mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("testing putput", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{Calls: 0}

		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go`

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}

		if spySleeper.Calls != 3 {
			t.Errorf("not enought calls to sleeper, want 3 got %d", spySleeper.Calls)
		}
	})

	t.Run("testing sequence", func(t *testing.T) {
		spySleepPrinter := &SpySleeperOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.SetDurationSlept}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("want %v but got %v", sleepTime, spyTime.durationSlept)
	}
}
