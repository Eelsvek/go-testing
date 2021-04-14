package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"motto": "yolo"}

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("hmm")

		assertError(t, err, ErrNotFound)
	})

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("motto")
		want := "yolo"

		assertStrings(t, got, want)
	})
}

func TestAdd(t *testing.T) {
	word := "cat"
	definition := "animal that meows"

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{word: definition}
		newDefinition := "animal that likes fish"

		err := dictionary.Add(word, newDefinition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func assertDefinition(t *testing.T, dictionary Dictionary, definition string, want string) {
	t.Helper()

	got, err := dictionary.Search("cat")
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != want {
		t.Errorf("want %q got %q", want, got)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertStrings(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}
