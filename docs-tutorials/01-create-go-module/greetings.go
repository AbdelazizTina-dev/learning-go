package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi Lord %v", name)
	return message
}