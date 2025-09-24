package main

import (
	"fmt"
)

func addNum(a int, b int) int {
	return a + b
}

func computeNum(a float64, b float64) (float64, float64){
	return a + b, a - b
}

func foo(a int, b int) (c int, d int){
	c = a * b
	d = a / b
	return
}

func main(){
	fmt.Println(addNum(10, 20))
	c, d := computeNum(2.13, 1.12)
	fmt.Println(c,d)
	e, f := foo(10, 3)
	fmt.Println(e, f)
}