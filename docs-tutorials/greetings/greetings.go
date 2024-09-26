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

func Hello_group(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func random_greeting() string {
	formats := []string{
		"Hi Lord %v",
		"Greetings lil %v",
		"Wassup %v",
	}

	return formats[rand.Intn(len(formats))]
}
