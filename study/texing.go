package main

import "fmt"

type Goods struct {
	Name  string
	Price int
}

type Book struct {
	Goods
	Writer string
}

func main() {
	b := Book{}
	b.Name = "asd"
	b.Price = 123
	b.Writer = "kljl"
	fmt.Println(b)
}
