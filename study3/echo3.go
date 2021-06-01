package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	startTime := time.Now().UnixNano()
	fd, _ := os.OpenFile("/dev/null", os.O_RDWR, 0666)
	for i := 0; i < 1000000; i++ {
		fmt.Fprintln(fd, strings.Join(os.Args[1:], " "))
	}
	defer fd.Close()

	endTime := time.Now().UnixNano()

	fmt.Println(endTime - startTime)
}
