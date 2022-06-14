package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(n, l string) string {
	if n == "" {
		n = "World"
	}

	prefix := englishHelloPrefix

	switch l {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	}

	return prefix + n
}

func main() {
	fmt.Println(Hello("World", "English"))
}
