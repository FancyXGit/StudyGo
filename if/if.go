package main

import (
	"fmt"
)

func main() {
	var a int;
	fmt.Scanf("%d", &a)
	if a > 10 {
		fmt.Println("a > 10")
	} else if a > 5 {
		fmt.Println("5 < a <= 10")
	} else {
		fmt.Println("a <= 5")
	}
}