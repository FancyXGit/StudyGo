package main

import (
	"fmt"
)

func main() {
	for a := 0; a < 10; a++{
		fmt.Println(a)
	}

	a := 1
	for a < 5{
		a++
		fmt.Println("NO WHILE")
	}
}