package main

import (
	"fmt"
	"log"
	"rsc.io/quote"
	"example.com/greetings"
)

func printQuote() {
	fmt.Println(quote.Go())
}

func greetPerson(name string) {
	message, err := greetings.Hello(name)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}

func main() {
	// Prepare Logger
	log.SetPrefix("greetings: ")
	log.SetFlags(0) // suppresses the time, source file, and line number

	// printQuote()
	greetPerson("Leo")
}
