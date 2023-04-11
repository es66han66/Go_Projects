package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(d))
	}

	if d[0] != "Spades of Ace" {
		t.Errorf("Expected 1st element to be Spades of Ace, but got %s", d[0])
	}
}

func TestSaveToFileAndGetDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()

	deck.saveToFile("_decktesting")

	loadedDeck := getDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in file, but got %d", len(deck))

	}

	os.Remove("_decktesting")
}
