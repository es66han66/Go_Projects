package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {
	name string
}
type spanishBot struct {
	name string
}

func (e englishBot) getGreeting() string {
	return e.name
}

func (s spanishBot) getGreeting() string {
	return s.name
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{
		name: "eshan",
	}

	sb := spanishBot{
		name: "eshanSpanish",
	}

	printGreeting(eb)

	printGreeting(sb)

}
