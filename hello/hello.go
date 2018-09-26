package main

import "fmt"

const french = "French"
const spanish = "Spanish"
const englishGreeting = "Hello, "
const frenchGreeting = "Bonjour, "
const spanishGreeting = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishGreeting
	case french:
		prefix = frenchGreeting
	default:
		prefix = englishGreeting
	}

	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
