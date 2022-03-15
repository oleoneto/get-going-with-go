package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"rsc.io/quote"
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

func greetPeople(names []string) {
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	for _, name := range names {
		fmt.Println(messages[name])
	}
}

func main() {
	// Prepare Logger
	log.SetPrefix("greetings: ")
	log.SetFlags(0) // suppresses the time, source file, and line number

	names := []string{"Leo", "John", "Oliver"}
	greetPeople(names)
}
