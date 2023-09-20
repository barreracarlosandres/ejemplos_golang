package main

import (
	"cards/calculator"
	"cards/channel"
	"cards/deck"
	"cards/http"
	"cards/interfaces"
	"cards/maps"
	"cards/mvc"
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
type calulatorExample struct{}

type examples interface {
	execute()
}

func main() {

	examples := map[examples]bool{
		deckExample{}:      false,
		interfaceExample{}: false,
		mapsExample{}:      false,
		httpExample{}:      false,
		structExample{}:    false,
		channelExample{}:   false,
		apiRestExample{}:   false,
		calulatorExample{}: true,
	}
	for example, shuldExecute := range examples {
		if shuldExecute {
			fmt.Println("**************************")
			example.execute()
		}
	}
}

func (ec calulatorExample) execute() {

	type Operacion interface {
		Calcular(d calculator.Datos, c chan string)
	}
	
	d := calculator.Datos{Dato1: 1, Dato2: 2}

	s := calculator.Suma{}
	r := calculator.Resta{}
	/*s.Calcular(d)
	r.Calcular(d)*/

	oprnes := []Operacion{s, r}

	c := make(chan string)

	for _, o := range oprnes {
		go o.Calcular(d, c)		
	}

	for i :=0; i < len(oprnes); i++ {
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
	mvc.RunApiRest()
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
