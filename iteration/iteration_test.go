package iteration

import (
	"fmt"
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
