package main

import "fmt"

func main() {
	var myArray1 [10]int

	myArray2 := [5]int{5,9,334,3244,9995}

	for i := 0; i < len(myArray1); i++{
		fmt.Printf("%d ", myArray1[i])
	}
	fmt.Printf("\n")

	for index, value := range myArray2 {
		fmt.Println("index = ", index, "value = ", value)
	}
}