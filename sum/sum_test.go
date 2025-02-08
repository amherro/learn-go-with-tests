package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Slice of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7}

		got := Sum(numbers)
		want := 28

		if got != want {
			t.Errorf("GOT %d WANT %d given %v", got, want, numbers)
		}
	})
}

// func TestSumAll(t *testing.T) {
// 	t.Run("takes in varying number of slices and returns the sum of each slice", func(t *testing.T) {
// 		got := SumAll([]int{1, 2}, []int{0, 9})
// 		want := []int{3, 9}

// 		if !reflect.DeepEqual(want, got) {
// 			t.Errorf("GOT %d WANT %d", got, want)
// 		}
// 	})
// }

func TestSumAllTails(t *testing.T) {
	testSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GOT %v WANT %v", got, want)
		}
	}
	t.Run("Sum all tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		testSums(t, got, want)
	})

	t.Run("Sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		testSums(t, got, want)
	})
}
