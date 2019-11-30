package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Lindsey")
		want := "Hello, Lindsey\n"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, World' when empty string given", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World\n"
		assertCorrectMessage(t, got, want)
	})
}
