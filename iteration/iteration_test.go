package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestIteration(t *testing.T) {
	t.Run("Repeat a character a specified number of times.", func(t *testing.T) {
		output := Iterate("F", 3)
		expected := "FFF"

		if output != expected {
			t.Errorf("Expected %q got %q", expected, output)
		}
	})
}

func BenchmarkIterate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iterate("F", 3)
	}
}

func ExampleIterate() {
	output := Iterate("F", 3)
	fmt.Println(output)
	// Output: FFF
}

// ************************************
// Testing different functions from the Go Strings package
// ************************************
func TestContains(t *testing.T) {
	got := strings.Contains("I love the Red Wings.", "Red Wings")
	expected := true

	if got != expected {
		t.Errorf("expected %t got %t", expected, got)
	}
}

func TestCount(t *testing.T) {
	// This is case-sensitive!! Putting in Adam fails ("Returns 1 not 2 b/c it doesn't count the uppercase A")
	got := strings.Count("adam", "a")
	expected := 2

	if got != expected {
		t.Errorf("expected %d got %d", expected, got)
	}
}

func TestPrefix(t *testing.T) {
	got1, got2 := strings.CutPrefix("Mr. Herro", "Mr")
	println(got1)
	want1 := ". Herro"
	want2 := true

	if got1 != want1 && got2 != want2 {
		t.Errorf("Want %q got %q", want1, got1)
		t.Errorf("Want %t got %t", want2, got2)
	}
}
