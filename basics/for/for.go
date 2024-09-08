package main

import "fmt"

func main() {
	for {
		fmt.Println("infinite")
		break
	}

	for i := range 3 {
		fmt.Println("range: ", i)
	}

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
}