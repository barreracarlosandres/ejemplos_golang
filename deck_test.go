package main

import "testing"

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
