package main

import (
	"cards/deck"
	"cards/http"
	"cards/interfaces"
	"cards/maps"
	"cards/structs"
	"fmt"
	"os"
)

type deckExample struct{}
type interfaceExample struct{}
type mapsExample struct{}
type httpExample struct{}
type structExample struct{}

type examples interface {
	execute()
}

func main() {

	examples := map[examples]bool{
		deckExample{} 		: false,
		interfaceExample{} 	: false,
		mapsExample{} 		: false,
		httpExample{} 		: true,
		structExample{} 	: false,
	}
	for e, b := range examples {
		if b {
			fmt.Println("**************************")
			e.execute()
		}
	}
}

func (he httpExample) execute(){
	http.GetHttp()
}

func newCard() string {
	return "Five of Diamon"
}

func (me mapsExample) execute() {
	fmt.Println("******* map example 1 *******")
	maps.Example1()
	fmt.Println("******* map example 2 *******")
	maps.Example2()
	fmt.Println("******* map example 3 *******")
	maps.Example3()
	fmt.Println("******* map example 4 *******")
	maps.Example4()
	fmt.Println("******* map example 5 *******")
	maps.Example5()
}

func (se structExample) execute() {

	//getPerson("Alex", "Perez")
	structs.GetPersonNew("Alex", "Perez")

}

func (e deckExample) execute() {
	cards :=   deck.NewDeck()

	hand, reamingCards := deck.Deal(cards, 6)
	hand.Print()
	reamingCards.Print()

	fmt.Println(cards.ToString())
	err := cards.ToSaveFile("my_cards")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1) //To exit de program
	}

	deck.NewDeckFromFile("my_cards").Print()

	cards.Shuffle()
	cards.Print()
}

func (i interfaceExample) execute() {
	eb := interfaces.EnglishBot{}
	sp := interfaces.SpanishBot{}

	interfaces.PrintGreeting(eb)
	interfaces.PrintGreeting(sp)
}
