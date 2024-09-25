package main

import (
	"fmt"
	"tutorial/first"
	"log"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("greetings: ")
	// Get a greeting message and print it.
	message, error := greetings.Hello("")

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(message)
}
