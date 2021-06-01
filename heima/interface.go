package main

import "fmt"

func main() {
	a := 123
	b := "asd"
	c := []string{"1", "asd"}

	f(a)
	f(b)
	f(c)
}

type A interface {
}

func f(a interface{}) {
	switch a.(type) {
	case bool:
		fmt.Println("bool")
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case []string:
		fmt.Println("[]string")
	}
}
