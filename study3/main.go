package main

import "fmt"

func main() {
	s := "撒娇快疯了asd"
	fmt.Println(string(reverse([]byte(s))))
}

func reverse(s []byte) []byte {
	srune := []rune(string(s))
	for i, j := 0, len(srune)-1; i < j; i, j = i+1, j-1 {
		srune[i], srune[j] = srune[j], srune[i]
	}
	return []byte(string(srune))
}
