package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const german = "German"
const czech = "Czech"
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const germanHelloPrefix = "Hallo, "
const czechHelloPrefix = "Ahoj, "
const englishHelloPrefix = "Hello, "

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case german:
		prefix = germanHelloPrefix
	case czech:
		prefix = czechHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name + "\n"
}

func main() {
	fmt.Printf(Hello("", ""))
}
