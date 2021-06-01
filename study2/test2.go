package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	var s1 string
	fmt.Print("输入第一段数据:")
	fmt.Scanln(&s)
	fmt.Print("输入第二段数据:")
	fmt.Scanln(&s1)
	q := strings.Split(s, ",")
	q1 := strings.Split(s1, ",")

	outFile, _ := os.OpenFile("xxx.csv", os.O_RDWR|os.O_CREATE, 0666)

	output := bufio.NewWriter(outFile)
	for i := 0; i < len(q); i++ {
		output.WriteString(q[i] + ",")
		output.WriteString(q1[i] + "\n")
	}
	output.Flush()
}
