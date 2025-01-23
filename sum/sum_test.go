package main

import "testing"

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
