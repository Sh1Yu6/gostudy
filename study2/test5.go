package main

import "fmt"

func f(a [3]int, b *[]int) {
	a[1] = 10
	*b = append(*b, 1, 2, 3, 4)
	fmt.Println(a)
	fmt.Println(b)
}
func main() {
	a := [3]int{1, 2, 3}
	b := []int{1, 2, 3}
	fmt.Println(a)
	fmt.Println(b)
	f(a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
