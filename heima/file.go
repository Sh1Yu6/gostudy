package main

import (
	"fmt"
	"os"

	"github.com/prometheus/common/log"
)

func main() {
	fd, err := os.Open("rand.go")
	if err != nil {
		log.Fatalln("open err", err)
	}
	fileInfo, err := fd.Stat()
	if err != nil {
		log.Fatalln("stat err", err)
	}
	fmt.Println(fileInfo.Name(), fileInfo.IsDir(), fileInfo.ModTime(),
		fileInfo.Size(), fileInfo.Mode())
}
