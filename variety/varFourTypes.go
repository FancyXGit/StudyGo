package main

import (
	"fmt"
)

func main(){
	var a int
	fmt.Printf("a = %d\n",a)
	fmt.Printf("Type of a = %T\n",a)

	var b int = 100
	fmt.Printf("b = %d\n",b)
	fmt.Printf("Type of b = %T\n",b)

	var c = 100
	fmt.Printf("c = %d\n",c)
	fmt.Printf("Type of c = %T\n",c)

	var d string = "GO!!!"
	fmt.Printf("d = %s\n",d)
	fmt.Printf("Type of d = %T\n",d)

	var e = "GOOOO!!!"
	fmt.Printf("e = %s\n",e)
	fmt.Printf("Type of e = %T\n",e)

	f := 3.1415
	fmt.Println("f = ",f)
	fmt.Printf("Type of f = %T\n",f)

	g := "yessss"
	fmt.Println("g = ",g)
	fmt.Printf("Type of g = %T\n",g)

	var i,j int = 100,200
	fmt.Println("i = ",i,",j = ",j)

	var k,l = 4.2141,"NOOOO"
	fmt.Println("k = ",k,",l = ",l)

	var (
		m int = 10
		n float64 = 3.1415
	)
	fmt.Println("m = ",m,",n = ",n)

	o,p := 4.342,10
	fmt.Println("o =",o,"p =",p)
}