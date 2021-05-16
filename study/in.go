package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scanln(&a)
	fmt.Println(a)

	var hens [10]int = [10]int{1, 3, 5, 7, 86, 4, 3, 2, 2, 3}
	fmt.Println(hens)

	for i := 0; i < 10; i++ {
		fmt.Println(hens[i])
	}

	for i, v := range hens[0:] {
		fmt.Println(i, " ", v)
	}

}
