package main

func main () {
	cards := newDeck()
	
	hand, reamingCards := deal(cards, 6)
	hand.print()
	reamingCards.print()
}

func newCard() string {
	return "Five of Diamon"
	
}