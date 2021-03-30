package main

const spanish = "Spanish"
const englishHello = "Hello, "
const spanishHello = "Hola, "

func Hello(name string, language string) string {
	var prefix = englishHello

	if language == spanish {
		prefix = spanishHello
	}

	if name == "" {
		name = "World"
	}

	return prefix + name
}
