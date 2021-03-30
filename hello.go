package main

const helloPrefix = "Hello, "
const defaultGreeting = "Hello, World"

func Hello(name string) string {
	if name == "" {
		return defaultGreeting
	}

	return helloPrefix + name
}
