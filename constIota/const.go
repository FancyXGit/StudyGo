package main

import (
	"fmt"
)

// const (
// 	BEIJING = 0
// 	SHANGHAI = 1
// 	SHENZHEN = 2
// )

const (
	BEIJING = iota * 10
	SHANGHAI
	SHENZHEN
)

const (
	GUANGZHOU = iota
)

func main(){
	const length int = 10
	fmt.Println("length = ",length)
	//length = 10
	fmt.Println(BEIJING,SHANGHAI,SHENZHEN,GUANGZHOU)
}