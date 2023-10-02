package main

import (
	"cards/apirest"
	"cards/apirestgin"
	"cards/calculator"
	"cards/channel"
	"cards/deck"
	"cards/http"
	"cards/interfaces"
	"cards/maps"
	"cards/pointers"
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
type apiRestExample struct{}
type calculatorExample struct{}
type pointerExample struct{}
type apiRestGinExample struct{}

type examples interface {
	execute()
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		message()
		return
	}

	switch example := args[0]; example {
	case "pointers":
		pointers.Example()
	case "structs":
		structs.Example()
	case "apiRestGin":
		apirestgin.Example()
	case "apiRest":
		apirest.Example()
	case "map":
		maps.Example()
	case "interfaces":
		interfaces.Example()
	default:
		message()
	}
}

func message() {
	fmt.Println("Ingresa una opción válida, ver README.MD")
}

func (pe pointerExample) execute() { pointers.Example() }

func (se structExample) execute() { structs.Example() }

func (me mapsExample) execute() { maps.Example() }

func (arge apiRestGinExample) execute() { apirestgin.Example() }

func (arge interfaceExample) execute() { interfaces.Example() }

func (apie apiRestExample) execute() {
	apirest.Example()
}

func (ec calculatorExample) execute() {

	type Operation interface {
		Calculator(d *calculator.Data, c chan string)
	}

	d := calculator.Data{First: 4, Second: 2}

	s := calculator.Add{}
	r := calculator.Subtraction{}
	/*s.Calculator(d)
	r.Calculator(d)*/

	oprnes := []Operation{s, r}

	c := make(chan string)

	for _, o := range oprnes {
		go o.Calculator(&d, c)
	}

	for i := 0; i < len(oprnes); i++ {
		fmt.Println(<-c)
	}

	/*go Operacion.Calculator(s, d, c)
	go Operacion.Calculator(r, d, c)*/

	/*for i := 0; i < 2; i++ {
		fmt.Println(<-c)
	}*/

	/*for l := range c {
		fmt.Println(l)
	}*/
}

func (ce channelExample) execute() {
	channel.Example1()
	channel.Example2()
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
