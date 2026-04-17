package main

import (
	"fmt"
)

const (
	englishHello = "Hello "
	spanishHello = "Hola "
	frenchHello  = "Bonjour "

	spanish = "Spanish"
	french  = "French"
)

func Hello(lang, name string) string {
	if name == "" {
		name = "World"
	}

	return generateGreetingLang(lang) + name
}
func generateGreetingLang(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishHello
	case french:
		prefix = frenchHello
	default:
		prefix = englishHello
	}
	return
}
func main() {
	fmt.Println(Hello("English", "Adam"))
}
