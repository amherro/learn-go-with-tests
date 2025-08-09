package iteration

import (
	"strings"
)

func Repeat(char string, iterations int) string {
	var repeated strings.Builder

	for i := 0; i < iterations; i++ {
		repeated.WriteString(char)
	}
	return repeated.String()
}

func PrintCharCount(str string, char string) int {
	return strings.Count(str, char)
}
