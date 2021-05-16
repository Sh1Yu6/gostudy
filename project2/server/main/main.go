package main

import (
	"fmt"
	"net"
	"time"

	"github.com/gostudy/project2/server/model"
)

func main() {

	initPool("127.0.0.1:6379", 16, 0, 300*time.Second)
	initUserDao()
	// 监听ip跟端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")

	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	for {
		// 等待连接
		conn, err := listen.Accept()

		if err != nil {
			fmt.Println("liste.Accept err:", err)
			continue
		}

		go process1(conn)
	}
}

func process1(conn net.Conn) {

	defer conn.Close()

	processor := &Processor{
		Conn: conn,
	}

	// 调用处理函数
	err := processor.process2()
	if err != nil {
		fmt.Println("processor err", err)
		return
	}
}

func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}
