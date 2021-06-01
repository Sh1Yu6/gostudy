package main

import "fmt"

type A interface {
	haha()
}

func doutai(a A) {
	a.haha()
}

type B struct {
}

func (t B) haha() {
	fmt.Println("bbbbbbbbbb")
}

type C struct {
}

func (t *C) haha() {
	fmt.Println("cccccccccc")
}

func main() {
	var b B
	var c C
	doutai(b)
	doutai(&c)
}
