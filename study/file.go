package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("main.go")

	if err != nil {
		fmt.Printf("open error: %v\n", err)
		os.Exit(1)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Printf("file close error: %v", err)
		}
	}(f)

	outFile, err := os.OpenFile("qwe.go", os.O_RDWR|os.O_APPEND|
		os.O_CREATE, 0666)

	if err != nil {
		fmt.Printf("open outFile error: %v\n")
		os.Exit(1)
	}
	defer outFile.Close()

	input := bufio.NewReader(f)
	output := bufio.NewWriter(outFile)

	for {
		str, err := input.ReadString('\n')
		if err != nil {
			fmt.Printf("input error: %v", err)
			break
		}
		output.WriteString(str)
	}
	output.Flush()
}
