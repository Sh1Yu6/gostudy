package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	start := time.Now().Unix()
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	//time.Sleep(1000 * time.Millisecond)

	str := "hello world"
	fmt.Println(strings.HasPrefix(str, "hello"))
	fmt.Println(strings.HasSuffix(str, "world"))
	//s := []rune(str)
	fmt.Println(len(str))
	fmt.Println(strings.Contains(str, "hell"))
	fmt.Println(strings.Count(str, "hell"))
	fmt.Println(strings.Index(str, "hell"))
	fmt.Println(strings.EqualFold(str, "hell"))

	end := time.Now().Unix()
	fmt.Println(end - start)
}
