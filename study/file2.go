package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type A struct {
	ChCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

func main() {
	inFile, _ := os.Open("main.go")
	defer inFile.Close()

	input := bufio.NewReader(inFile)

	var res A
	for {
		str, err := input.ReadString('\n')
		if err == io.EOF {
			break
		}
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z':
				res.ChCount++
			case v >= '0' && v <= '9':
				res.NumCount++
			case v == ' ' || v == '\t' || v == '\n':
				res.SpaceCount++
			default:
				res.OtherCount++
			}
		}
	}
	fmt.Printf("c=%d, i=%d, s=%d, o=%d", res.ChCount,
		res.NumCount, res.SpaceCount, res.OtherCount)

}
