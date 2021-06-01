package main

import "fmt"

func f1(a [5]int) {
	fmt.Println(&a)
}

func f2(b []int) {
	fmt.Println(&b)
}
func main() {
	a := [5]int{}
	b := []int{}

	fmt.Println(a)
	fmt.Println(b)
	f1(a)
	f2(b)
}
