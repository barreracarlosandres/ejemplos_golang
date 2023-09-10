package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	var dl = 16
	if len(d) != dl {
		t.Errorf("Expected deck lenght of %v, but got: %v", dl, len(d))
	}

	var ds = "Ace of Spades"
	if d[0] != ds {
		t.Errorf("Expected first card %v, but got %v:", ds, d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.toSaveFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
