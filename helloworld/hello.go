package main

import "fmt"

const (
	french        = "French"
	spanish       = "Spanish"
	englishPrefix = "Hello, "
	frenchPrefix  = "Bonjour, "
	spanishPrefix = "Hola, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return selectPrefix(language) + name + "!"
}

func selectPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Danil", ""))
}
