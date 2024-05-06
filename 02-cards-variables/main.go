package main

import "fmt"

// someVar := "Some Value" // This is not valid outside of a function

// Note: Variables must first be initialized with the ':=' operator or the 'var variableName type' syntax

// var someVar string

func main() {
	// var card string = "Ace of Spades"

	// This is a shorthand way of declaring a variable, assigning a value to it, and assigning a type to it
	// Only have to use this shorthand when declaring a new variable (not when reassigning a value to an existing variable)
	// card := "Ace of Spades"

	// This is a way of declaring a variable
	// var deckSize int
	// And then assigning a value to it
	// deckSize = 52

	// someVar = "Some Value"

	card := newCard()

	fmt.Println(card)
}

// This is a function returns a string, and explicitly typing out "string" tells the go compiler that this function will return a type string
func newCard() string {
	return "Five of Diamonds"
}
