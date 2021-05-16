package main

import "fmt"

type A struct {
	a int
	b int
	c string
}

type Test struct {
	a int
	b int
	c string
}

func (a A) test() {
	fmt.Println("hasdfjkl")
}

func (a *Test) test() {
	a.a = 100
	fmt.Println("aaaaaaaaaaaa")
}

func (a *Test) String() string {
	return "jsdakfjaskldfjkl"
}

func main() {

	obj1 := A{1, 2, "asd"}
	obj2 := new(A)
	obj3 := &A{}
	fmt.Println(obj1.a)
	fmt.Println(obj2.b)
	fmt.Println(obj3.a)
	obj1.test()
	obj2.test()
	obj3.test()

	obj4 := Test{}
	obj4.test()
	fmt.Println(obj4)

}
