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

		examples := map[examples]bool{
			pointerExample{}:    true,
			structExample{}:     true,
			mapsExample{}:       true,
			deckExample{}:       false,
			interfaceExample{}:  false,
			httpExample{}:       false,
			channelExample{}:    false,
			apiRestExample{}:    false,
			calculatorExample{}: false,
			apiRestGinExample{}: false,
		}
		for example, shouldExecute := range examples {
			if shouldExecute {
				example.execute()
			}
		}
	} else if args[0] == "pointers" {
		pointers.Example()
	} else if args[0] == "apiRest" {
		apirest.RunApiRest()
	} else if args[0] == "apiRestGin" {
		apirestgin.Example()
	} else {
		fmt.Println("Las opciones son:\n" +
			"pointers\n" +
			"maps",
		)
	}
}

func (pe pointerExample) execute() { pointers.Example() }

func (se structExample) execute() { structs.Example() }

func (me mapsExample) execute()         { maps.Example() }
func (arge apiRestGinExample) execute() { apirestgin.Example() }

func (ec calculatorExample) execute() {

	type Operation interface {
		Calcular(d *calculator.Datos, c chan string)
	}

	d := calculator.Datos{Dato1: 4, Dato2: 2}

	s := calculator.Suma{}
	r := calculator.Resta{}
	/*s.Calcular(d)
	r.Calcular(d)*/

	oprnes := []Operation{s, r}

	c := make(chan string)

	for _, o := range oprnes {
		go o.Calcular(&d, c)
	}

	for i := 0; i < len(oprnes); i++ {
		fmt.Println(<-c)
	}

	/*go Operacion.Calcular(s, d, c)
	go Operacion.Calcular(r, d, c)*/

	/*for i := 0; i < 2; i++ {
		fmt.Println(<-c)
	}*/

	/*for l := range c {
		fmt.Println(l)
	}*/
}

func (apie apiRestExample) execute() {
	apirest.RunApiRest()
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

func (i interfaceExample) execute() {
	eb := interfaces.EnglishBot{}
	sp := interfaces.SpanishBot{}

	interfaces.PrintGreeting(eb)
	interfaces.PrintGreeting(sp)
}
