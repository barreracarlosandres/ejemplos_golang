package interfaces

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func (eb englishBot) getGreeting() string {
	return "englishBot"
}

func (eb spanishBot) getGreeting() string {
	return "spanishBot"
}

func printGreeting(e bot) {
	fmt.Println(e.getGreeting())
}

func Example() {
	fmt.Println("------------ Interfaces Examples ------------")
	eb := englishBot{}
	sp := spanishBot{}

	printGreeting(eb)
	printGreeting(sp)
}
