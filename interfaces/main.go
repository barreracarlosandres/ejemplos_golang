package interfaces

import "fmt"

type bot interface {
	getGreeting() string
}

type EnglishBot struct{}
type SpanishBot struct{}

func (eb EnglishBot) getGreeting() string {
	return "englishBot"
}

func (eb SpanishBot) getGreeting() string {
	return "spanishBot"
}

func PrintGreeting(e bot) {
	fmt.Println(e.getGreeting())
}