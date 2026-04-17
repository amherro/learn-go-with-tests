package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	got := Iteration("a", 5)
	want := "aaaaa"

	if got != want {
		t.Errorf("GOT %q want %q", got, want)
	}
}

func BenchmarkIteration(b *testing.B) {
	for b.Loop() {
		Iteration("a", 5)
	}
}

func ExampleIteration() {
	letters := Iteration("a", 7)
	fmt.Println(letters)
	// Output: aaaaaaa
}
