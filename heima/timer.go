package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)
	fmt.Println(time.Now())
	c := timer.C
	fmt.Println(<-c)
	fmt.Println(<-time.After(time.Second * 3))

}
