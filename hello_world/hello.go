package main

import "fmt"

const (
	spanish = "Spanish"
	french = "French"
	serbian = "Serbian"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
	serbianHelloPrefix = "Zdravo, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case serbian:
		prefix = serbianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

func main(){
	fmt.Println(Hello("Elodie", "Spanish"))
}