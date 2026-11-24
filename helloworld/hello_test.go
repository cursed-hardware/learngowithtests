package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello by name", func(t *testing.T) {
		got := Hello("Danil", "")
		want := "Hello, Danil!"
		assertCorrMsg(t, got, want)
	})

	t.Run("saying 'Hello, World!' when empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrMsg(t, got, want)
	})

	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Danil", "French")
		want := "Bonjour, Danil!"
		assertCorrMsg(t, got, want)
	})

	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Danil", "Spanish")
		want := "Hola, Danil!"
		assertCorrMsg(t, got, want)
	})
}

func assertCorrMsg(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
