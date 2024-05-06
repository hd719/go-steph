package main

import "fmt"

func main() {
	// Note: Both arrays and slices have fixed lengths, but slices are more flexible than arrays
	// and a slice can only have 1 type of data in it
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades") // does not mutate the existing slice, but instead returns a new slice

	// Re-initialize the i and card variables on each iteration of the loop, so we have to use the ':=' operator
	for i, card := range cards {
		fmt.Println(i, card)
	}

	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
