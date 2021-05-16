package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3}
	b := [...]int{1, 2, 3}

	fmt.Println(a == b)
}
