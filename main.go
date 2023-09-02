package main

import "fmt"

func main () {
	card := newCard()
	fmt.Println(card)
	cards := []string{newCard(), newCard()}
	cards = append(cards, "Otro datos")
	fmt.Println(cards)

	//Recorrer all records of cards slice (array)
	for i, card := range cards {
		fmt.Println(i, card)	
	}
}