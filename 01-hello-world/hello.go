package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(n, l string) string {
	if n == "" {
		n = "World"
	}

	var prefix string

	switch l {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix + n
}

func main() {

	name := "Ryan"

	fmt.Println(Hello(name, "English"))
}
