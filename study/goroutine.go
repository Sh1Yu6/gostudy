package main

import (
	"fmt"
	"sync"
	"time"
)

var N int
var lock sync.Mutex

func test() {
	for i := 0; i < 100000; i++ {
		lock.Lock()
		N++
		lock.Unlock()
	}
}

func test2() {
	for i := 0; i < 100000; i++ {
		lock.Lock()
		N++
		lock.Unlock()
	}
}
func main() {
	go test()
	go test()
	time.Sleep(time.Second * 5)
	lock.Lock()
	fmt.Println(N)
	lock.Unlock()
}
