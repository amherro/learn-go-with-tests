package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to the inputted name", func(t *testing.T) {
		got := Hello("Adam", "")
		want := "Hello Adam"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say \"Hello World\" when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Adam", "Spanish")
		want := "Hola Adam"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Adam", "French")
		want := "Bonjour Adam"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Dutch", func(t *testing.T) {
		got := Hello("Adam", "Dutch")
		want := "Hallo Adam"

		assertCorrectMessage(t, got, want)
	})
}
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
