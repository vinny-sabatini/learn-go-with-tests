package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("saying hello to someone", func(t *testing.T) {
		got := Hello("Vinny", "")
		want := "Hello, Vinny"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to no one", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in Spanish", func(t *testing.T) {
		got := Hello("Amigo", "Spanish")
		want := "Hola, Amigo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in French", func(t *testing.T) {
		got := Hello("Monsieur", "French")
		want := "Bonjour, Monsieur"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in German", func(t *testing.T) {
		got := Hello("Herr", "German")
		want := "Hallo, Herr"
		assertCorrectMessage(t, got, want)
	})
}
