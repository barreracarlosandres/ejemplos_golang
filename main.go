package main

func main () {
	//cards := deck{newCard(), newCard()}
	//cards = append(cards, "Otro datos")
	//fmt.Println(cards)

	cards := newDeck()
	cards.print()
}

func newCard() string {
	return "Five of Diamon"
	
}