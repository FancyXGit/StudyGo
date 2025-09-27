package main

import "fmt"

func main() {
	numbers := []int{32, 34, 32, 5, 4, 6, 3442, 23, 34, 321}
	fmt.Printf("len = %d, cap = %d, vec = %v\n", len(numbers), cap(numbers), numbers)
	numbers = append(numbers, 12)
	fmt.Printf("len = %d, cap = %d, vec = %v\n", len(numbers), cap(numbers), numbers)

	numbers1 := numbers[0:2] //numbers[2] not included
	fmt.Printf("len = %d, cap = %d, vec = %v\n", len(numbers1), cap(numbers1), numbers1)

	numbers1[0] = 1 //share same address
	fmt.Printf("numbers: len = %d, cap = %d, vec = %v\n", len(numbers), cap(numbers), numbers)
	fmt.Printf("numbers1: len = %d, cap = %d, vec = %v\n", len(numbers1), cap(numbers1), numbers1)

	numbers2 := make([]int, 3)
	copy(numbers2, numbers[3:6]) //copy value
	fmt.Printf("len = %d, cap = %d, vec = %v\n", len(numbers2), cap(numbers2), numbers2)
}
