package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("english", "Adam")
		want := "Hello Adam"

		assertCorrectMessage(t, got, want)

	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"

		assertCorrectMessage(t, got, want)

	})
	t.Run("say hello in Spanish", func(t *testing.T) {
		got := Hello("Spanish", "Adam")
		want := "Hola Adam"

		assertCorrectMessage(t, got, want)

	})
	t.Run("say hello in French", func(t *testing.T) {
		got := Hello("French", "Adam")
		want := "Bonjour Adam"

		assertCorrectMessage(t, got, want)

	})
}
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
