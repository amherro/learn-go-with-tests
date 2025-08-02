package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say hello to people", func(t *testing.T) {
		got := Hello("Adam", "English")
		want := "Hello Adam"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Print 'Hello World' if an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Passing in Spanish", func(t *testing.T) {
		got := Hello("Adam", "Spanish")
		want := "Hola Adam"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Pass in French", func(t *testing.T) {
		got := Hello("Adam", "French")
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
