package main

import "fmt"

func main() {
	var m map[int]string = make(map[int]string)
	m[1] = "asdj"
	fmt.Println(m[1])
	m[1] = "asd"
	var s []int
	s = append(s, 1)
	fmt.Println(s[0])

	var s []map[int]string
}
