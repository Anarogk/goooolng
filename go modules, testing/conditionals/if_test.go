package main

import "testing"

func TestIf(t *testing.T) {
	// subtests
	t.Run("saying hello to ppl", func(t *testing.T) {
		got := Hello("Chris", "Eng")
		want := "Hello! Chris"
		assertCorrectMessage(t, got, want)

	})
	// subtests
	t.Run("Say \"Hello! World\" when empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello! World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("check language specific functionality", func(t *testing.T) {
		got := Hello("Asta", "Spanish")
		want := "Hola! Asta"
		assertCorrectMessage(t, got, want)
	})
}

// refactored assertion into function
func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
