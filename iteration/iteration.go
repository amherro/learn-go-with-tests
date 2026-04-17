package iteration

import (
	"fmt"
	"strings"
)

func Iteration(letter string, iterations int) string {
	var repeated strings.Builder
	for i := range iterations {
		repeated.WriteString(letter)
		i++
	}
	return repeated.String()
}

// Implementing FizzBuzz test
func FizzBuzz() {
	for i := range 100 {
		i++
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		} else if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		} else if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		fmt.Println(i)
	}
}
