package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got: %q want: %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Ryan", "English")
		want := "Hello, Ryan"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Frend", func(t *testing.T) {
		got := Hello("Jean", "French")
		want := "Bonjour, Jean"

		assertCorrectMessage(t, got, want)
	})

	t.Run("unknown language defualt to English", func(t *testing.T) {
		got := Hello("Nagate", "")
		want := "Hello, Nagate"

		assertCorrectMessage(t, got, want)
	})
}
