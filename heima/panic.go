package main

import "fmt"

func main() {
	defer func() {
		recover()
		fmt.Println("a")
	}()
	f()
	fmt.Println("qqq")
}

func f() {
	//defer func() {
	//recover()

	//}()
	panic("aksdjaklsdjkl")
}
