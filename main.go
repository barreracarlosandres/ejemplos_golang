package main

import (
	"fmt"
	"os"
)

/*import (
	"fmt"
	"os"
)*/

func main() {

	//sliceExamples()
	//structExamples()
	mapExamples()
}

func newCard() string {
	return "Five of Diamon"
}

func mapExamples() {
	fmt.Println("******* map example 1 *******")
	mapExample1()
	fmt.Println("******* map example 2 *******")
	mapExample2()
	fmt.Println("******* map example 3 *******")
	mapExample3()
	fmt.Println("******* map example 4 *******")
	mapExample4()
	fmt.Println("******* map example 5 *******")
	mapExample5()
}

func structExamples() {

	//getPerson("Alex", "Perez")
	getPersonNew("Alex", "Perez")

}

func sliceExamples() {
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

	cards.shuffle()
	cards.print()
}
