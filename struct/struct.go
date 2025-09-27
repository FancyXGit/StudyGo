package main

import "fmt"

type Book struct {
	title  string
	author string
	price  float64
}

func printBook(myBook Book) {
	fmt.Println("title:", myBook.title)
	fmt.Println("author:", myBook.author)
	fmt.Println("price:", myBook.price)
}

func main() {
	var Book1 Book
	Book1.author = "ZhangSan"
	Book1.title = "Go Language Programming"
	Book1.price = 120
	printBook(Book1)
}
