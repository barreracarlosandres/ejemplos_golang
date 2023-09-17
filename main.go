package main

import (
	"cards/channel"
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
type channelExample struct{}

type examples interface {
	execute()
}

func main() {

	examples := map[examples]bool{
		deckExample{}		: false,
		interfaceExample{}	: false,
		mapsExample{}		: false,
		httpExample{}		: false,
		structExample{}		: false,
		channelExample{}	: true,
	}
	for example, shuldExecute := range examples {
		if shuldExecute {
			fmt.Println("**************************")
			example.execute()
		}
	}
}

func (ce channelExample) execute() {
	channel.Example1()
}

func (he httpExample) execute() {
	fmt.Println("###### Examples of interfaces #######")
	//http.GetHttpExample1("http://google.com")
	//http.GetHttpExample2("http://google.com")
	http.GetHttpExample3("http://google.com")
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
	cards := deck.NewDeck()

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
