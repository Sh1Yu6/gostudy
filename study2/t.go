package main

import "fmt"

func f() (int, error) {
	return 0, nil
}
func main() {
	if a, b := f(); b != nil {
		fmt.Println("asd")
	} else {
		fmt.Println(a)
	}
	var c int
	e := &c
	f := &e

	fmt.Println(&c)
	fmt.Println(&e)
	fmt.Println(**f)

	s := make([]int, 2, 10)
	s1 := []int{1, 2, 3, 4}

	copy(s, s1)
	fmt.Println(s)
	s = append(s, 1, 2, 3, 4, 5)
	fmt.Println(s)
}
