package main

import (
	"fmt"
)

const (
	englishHelloPrefix string = "Hello "
	spanishHelloPrefix string = "Hola "
	frenchHelloPrefix  string = "Bonjour "
	dutchHelloPrefix   string = "Hallo "

	spanish string = "Spanish"
	french  string = "French"
	dutch   string = "Dutch"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return chooseLanguage(language) + name
}

func chooseLanguage(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case dutch:
		prefix = dutchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Adam", "Dutch"))
}
