package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 19)
	want := "aaaaaaaaaaaaaaaaaaa"

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}

func TestCount(t *testing.T) {
	got := PrintCharCount("Hello", "l")
	want := 2

	if got != want {
		t.Errorf("Got %d expected %d", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 19)
	}
}

func ExampleRepeat() {
	repeated := Repeat("b", 5)
	fmt.Println(repeated)
	// Output: bbbbb
}
