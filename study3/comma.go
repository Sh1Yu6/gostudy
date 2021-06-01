package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var s string
	fmt.Scanln(&s)
	t := s
	n := len(s)
	if n <= 3 {
		//return s
		fmt.Println(s)
		os.Exit(0)
	}
	s = s[:n-3] + "," + s[n-3:]
	fmt.Println(s)
	fmt.Println(comm(t))
}

func comm(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		if i > 0 && i%3 == 0 {
			buf.WriteString(",")
			buf.WriteByte(s[i])
		} else {
			buf.WriteByte(s[i])
		}
	}
	return buf.String()
}
