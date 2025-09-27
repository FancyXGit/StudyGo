package main

import "fmt"

func deferFunc() int {
	fmt.Println("DEFER FUNC")
	return 0
}

func returnFunc() int {
	fmt.Println("RETURN")
	return 0
}

func test01() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	fmt.Println("123")

	defer fmt.Println("DEFER")

	fmt.Println("456")

	test01()
}