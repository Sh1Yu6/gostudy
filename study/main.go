package main

import (
	"flag"
	"fmt"
	"strings"
)

func Addupper() func(int) int {
	n := 10
	return func(x int) int {
		n += x
		return n
	}
}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) == false {
			return name + suffix
		}
		return name
	}
}

func main() {
	//fmt.Println("len=", len(os.Args), " s=", os.Args)

	//for _, v := range os.Args {
	//fmt.Println(v)
	//}
	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user, "u", "", "获取用户名, 默认为空")
	flag.StringVar(&pwd, "pwd", "", "获取mima, 默认为空")
	flag.StringVar(&host, "h", "localhost", "获取host, 默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号, 默认3306")
	flag.Parse()

	fmt.Printf("user=%v, pwd=%v, host=%v, port=%v\n",
		user, pwd, host, port)
}
