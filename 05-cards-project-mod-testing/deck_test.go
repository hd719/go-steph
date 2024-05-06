package main

import (
	"os"
	"testing"
)

// t *testing.T is a type of variable
// t represents the test handler and is the bridge between the test code and the testing package
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		// %v is a placeholder for the value of len(d)
		// t.Errorf() is a function that prints out an error message and causes the test to fail
		// Notifying the the test handler (t) that something went wrong
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but go %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// Remove any files that may have been created by previous tests
	os.Remove("_decktesting")

	// Create a new deck
	deck := newDeck()

	// Save the deck to a file
	deck.saveToFile("_decktesting")

	// Load the deck from the file
	loadDeck := newDeckFromFile("_decktesting")

	if len(loadDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadDeck))
	}

	// Remove the file
	os.Remove("_decktesting")
}
