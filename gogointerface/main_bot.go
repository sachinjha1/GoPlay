package main

import "fmt"

type englishBot struct{}

func (englishBot) getGreeting() string {
	//very custom englishBot logic
	return "Hello"
}

type spanishBot struct{}

func (spanishBot) getGreeting() string {
	//very custom logic of spanishbot
	return "Hola"
}

/*
  Concrete Type -> map, struct, int, string, englishBot
  Interface Type -> bot
*/
//Any struct which has implementation of getGreeting() function which returns string
//will get automatically associated with bot.
type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
