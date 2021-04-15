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
	t.Run("new word", func(t *testing.T) {
		word := "cat"
		definition := "animal that meows"

		dictionary := Dictionary{}
		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "cat"
		definition := "animal that meows"

		dictionary := Dictionary{word: definition}
		newDefinition := "animal that likes fish"

		err := dictionary.Add(word, newDefinition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestDelete(t *testing.T) {
	word := "dog"
	definition := "animal that barks"
	dictionary := Dictionary{word: definition}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)

	if err != ErrNotFound {
		t.Errorf("wanted error %q got %q", ErrNotFound, err)
	}
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "dog"
		definition := "animal that barks"

		dictionary := Dictionary{word: definition}
		newDefinition := "animal that is inferior to cat"

		dictionary.Update(word, newDefinition)

		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}

		word := "dog"
		definition := "animal that is inferior to cat"

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrorWordDoesNotExist)
	})
}

func assertDefinition(t *testing.T, dictionary Dictionary, word string, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != definition {
		t.Errorf("got %q want %q", got, definition)
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
