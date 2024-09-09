package main

import (
	"fmt"
)

func main() {
	var a[5]int 
	fmt.Println(a)

	a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])


	b := [...]int{100, 3: 400, 500}
    fmt.Println("idx:", b)
}