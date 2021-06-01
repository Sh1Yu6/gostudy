package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fd, err := os.OpenFile("rand.go", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("openfile err", err)
	}
	buf := make([]byte, 6)
	n, err := fd.Read(buf)
	fmt.Println(string(buf[:n]))
	n, err = fd.Read(buf)
	fmt.Println(string(buf[:n]))
	fmt.Println(buf)

	a := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", a[i][j])
		}
		fmt.Println()
	}
}
