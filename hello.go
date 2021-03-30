package main

const spanish = "Spanish"
const french = "French"
const englishHello = "Hello, "
const spanishHello = "Hola, "
const frenchHello = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHello
	case spanish:
		prefix = spanishHello
	default:
		prefix = englishHello
	}

	return
}
