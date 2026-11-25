package maps

import (
	"errors"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "this is just a test"

		assertError(t, err, nil)
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrUnknownWord

		assertError(t, err, ErrUnknownWord)
		assertStrings(t, err.Error(), want.Error())
	})
}

func TestAdd(t *testing.T) {
	t.Run("add new word to dictionary", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "this is just a test")

		want := "this is just a test"
		got, err := dictionary.Search("test")

		assertError(t, err, nil)
		assertStrings(t, got, want)
	})

	t.Run("redefinition of existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new definiton")
		assertError(t, err, ErrRedefineWord)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)
		updatedDefinition, _ := dictionary.Search(word)

		assertError(t, err, nil)
		assertStrings(t, newDefinition, updatedDefinition)
	})

	t.Run("update non existing word", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Update("test", "this is just a test")

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "this is just a test"}

		err := dictionary.Delete(word)
		assertError(t, err, nil)

		_, err = dictionary.Search(word)
		assertError(t, err, ErrUnknownWord)
	})

	t.Run("delete non existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}

		err := dictionary.Delete(word)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertError(t testing.TB, err, expectedErr error) {
	t.Helper()
	if !errors.Is(err, expectedErr) {
		t.Fatalf("got %q but expected %q", err.Error(), expectedErr.Error())
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given", got, want)
	}
}
