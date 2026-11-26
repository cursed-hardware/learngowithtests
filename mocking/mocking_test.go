package mocking

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
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
}
