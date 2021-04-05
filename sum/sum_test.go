package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %q want %q given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{10, 9})
	want := []int{3, 19}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSum := func(t *testing.T, got []int, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("sums the tails of provided slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{10, 10, 10})
		want := []int{2, 20}

		checkSum(t, got, want)

	})

	t.Run("sums empty slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{})
		want := []int{2, 0}

		checkSum(t, got, want)
	})
}
