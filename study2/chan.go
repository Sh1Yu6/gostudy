package main

import (
	"fmt"
	"time"
)

func main() {
	nBufChan := make(chan int, 10)
	closeChan := make(chan bool)
	go wr(nBufChan)
	//for i := range nBufChan {
	go re(nBufChan, closeChan)
	//}
	<-closeChan
	fmt.Println("close")
}

func wr(nBufChan chan int) {
	for i := 0; i < 100; i++ {
		nBufChan <- i
		time.Sleep(time.Millisecond * 500)
	}
	close(nBufChan)
}

func re(nBufChan chan int, closeChan chan bool) {
	for n := range nBufChan {
		fmt.Println(n)
	}
	closeChan <- true
	close(closeChan)
}
