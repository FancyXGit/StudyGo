package main

import "fmt"

func printMap(myMap map[string]string) {
	for key, value := range myMap {
		fmt.Printf("key = %s, value = %s\n", key, value)
	}
}

func main() {
	var myMap1 map[string]int
	fmt.Println("myMap1", myMap1 == nil)

	myMap2 := make(map[string]int, 5)
	fmt.Println("myMap2", myMap2 == nil)
	myMap2["C++"] = 1
	myMap2["Python"] = 2
	myMap2["Java"] = 3
	myMap2["Go"] = 2147483647
	fmt.Println(myMap2)

	myMap3 := map[string]string{
		"one":   "php",
		"two":   "java",
		"three": "python",
	}
	fmt.Println(myMap3)
	printMap(myMap3)
	delete(myMap3, "two")
	printMap(myMap3)
}
