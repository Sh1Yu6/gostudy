package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	fmt.Print("输入第一段数据:")
	fmt.Scanln(&s)
	q := strings.Split(s, ",")

	outFile, _ := os.OpenFile("xxxx.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	output := bufio.NewWriter(outFile)
	for i := 0; i < len(q); i++ {
		output.WriteString(q[i] + "\n")
	}
	output.Flush()
}
