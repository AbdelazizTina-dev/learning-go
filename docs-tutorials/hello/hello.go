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
	messages, error := greetings.Hello_group([]string{"Peter", "Ned", "Cat"})

	if error != nil {
		log.Fatal(error)
	}

	for _, message := range messages {
		fmt.Println(message)
	}
}
