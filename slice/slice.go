package main

import "fmt"

func printArr(arr []int) {
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}

func main() {
	myArray := []int{45, 346, 34, 12}
	printArr(myArray)

	slice1 := []int{23,34,21}
	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)

	slice2 := make([]int, 3)
	fmt.Printf("len = %d, slice = %v\n", len(slice2), slice2)

	slice3 := []int{}
	fmt.Println(slice3 == nil)

	var slice4 []int
	fmt.Println(slice4 == nil)
}