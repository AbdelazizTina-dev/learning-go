package main

import (
	"fmt"
	"log"
	"tutorial/first"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("greetings: ")
	// Get a greeting message and print it.
	message, error := greetings.Hello("Peter")

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(message)
}
