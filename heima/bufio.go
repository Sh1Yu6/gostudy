package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	buf := make([]byte, 20)
	//input.Read(buf)
	s, _ := input.ReadSlice('\n')
	fmt.Println([]byte(s))
	output.Write(buf)

	ss := bufio.NewScanner(os.Stdin)
	for ss.Scan() {
		fmt.Println(ss.Text())
		fmt.Println("a")
	}
}
