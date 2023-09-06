package main

import (
	"fmt"
	"os"
)

func main() {
	cards := newDeck()

	hand, reamingCards := deal(cards, 6)
	hand.print()
	reamingCards.print()

	fmt.Println(cards.toString())
	err := cards.toSaveFile("my_cards")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1) //To exit de program
	}

	newDeckFromFile("my_cards").print()
}

func newCard() string {
	return "Five of Diamon"
}
