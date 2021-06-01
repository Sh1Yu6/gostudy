package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	fmt.Println(t1)
	t2 := time.Now().Format("2006年1月2日 15:04:05")
	fmt.Println(t2)

	year, month, day := time.Now().Date()
	fmt.Println(year, month, day)
	fmt.Println(time.Now().Date())

}
