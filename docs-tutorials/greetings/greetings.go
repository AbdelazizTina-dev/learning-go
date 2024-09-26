package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(random_greeting(), name)
	return message, nil
}

func random_greeting() string {
	formats := []string{
		"Hi Lord %v",
		"Greetings lil %v",
		"Wassup %v",
	}

	return formats[rand.Intn(len(formats))]
}
