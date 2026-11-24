package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Danil")
	want := "Hello, Danil!"

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
