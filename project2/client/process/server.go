package process

import (
	"fmt"
	"net"

	"github.com/gostudy/project2/common/utils"
)

func ShowMenu() {
	for {
		utils.CallClear()
		fmt.Println("----------------聊天系统---------------")
		fmt.Println("1) 显示在线用户列表")
		fmt.Println("2) 发送消息")
		fmt.Println("3) 信息列表")
		fmt.Println("4) 退出系统")
		fmt.Print("请选择(0-3):")

		var key int
		fmt.Scanln(&key)
		switch key {
		case 1:
		case 2:
		case 3:
		case 0:
			fmt.Println("退出")
			return
		default:
			fmt.Println("输入错误, 请重新输入!")
		}

	}
}

func serverProcessMsg(conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		msg, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("serverProcessMsg error", err)
			return
		}
		fmt.Println(msg)
	}
}
