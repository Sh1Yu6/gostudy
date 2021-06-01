package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	start := time.Now()
	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Int())
	fmt.Println(rand.Int())
	fmt.Println(time.Since(start))

	f1 := a(10)
	f2 := a(100)
	fmt.Println(f1())
	fmt.Println(f2())

	s := []int{5, 43, 2, 3, 1, 100}
	sort.Ints(s)
	fmt.Println(s)

}

func a(x int) func() int {
	value := x
	return func() int {
		var res int
		for i := 1; i <= value; i++ {
			res += i
		}
		return res
	}
}
