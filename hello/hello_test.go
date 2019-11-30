package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Lindsey")
	want := "Hello, Lindsey\n"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
