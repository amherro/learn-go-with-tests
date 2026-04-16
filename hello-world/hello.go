package main

import "fmt"

const englishHello = "Hello "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHello + name
}

func main() {
	fmt.Println(Hello("Adam"))
}
