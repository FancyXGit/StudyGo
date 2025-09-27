package main

import (
	"fmt"
)

func main() {
	for a := range 10{
		fmt.Println(a)
	}

	a := 1
	for a < 5{
		a++
		fmt.Println("NO WHILE")
	}
}