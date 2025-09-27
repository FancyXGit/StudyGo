package main

import (
	"fmt"
)

func mySwap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	a := 10
	b := 20
	mySwap(&a, &b)
	fmt.Println(a, b)
}