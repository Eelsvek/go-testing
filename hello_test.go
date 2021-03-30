package main

import "testing"

func TestHello(t *testing.T) {
	assertion := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Kevin")
		want := "Hello, Kevin"

		assertion(t, got, want)
	})

	t.Run("saying hello world if empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertion(t, got, want)
	})
}
