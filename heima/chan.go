package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int, 100)
	go a(c)
	//go a(c)
	//go b(c)
	for v := range c {
		time.Sleep(time.Second)
		fmt.Println("close", v)
	}
}

func a(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
	fmt.Println("kj")
}

//func b(c chan bool) {
//for i := 0; i < 10; i++ {
//time.Sleep(time.Second)
//fmt.Println("aaaa", i)
//}
//c <- true
//}
