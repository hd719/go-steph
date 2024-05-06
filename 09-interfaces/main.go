package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

// This interface is being used by both englishBot and spanishBot and if you are type in this program with a function called getGreeting() that returns a string, you are now an honorary member of type bot
// And because you are an honorary member of type bot, you can now call printGreeting() with you as an argument
// This interface tells us what methods/functions can be called on a type
// With interfaces you cannot create values with them, you cannot say var b bot = 5
type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// Interfaces

func (englishBot) getGreeting() string {
	// Very custom logic for generating an english greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	// Very custom logic for generating a spanish greeting
	return "Hola!"
}
