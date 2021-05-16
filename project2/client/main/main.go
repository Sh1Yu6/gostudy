package main

import (
	"fmt"
	"os"

	"github.com/gostudy/project2/client/process"
	"github.com/gostudy/project2/common/utils"
)

func main() {
	start()
}

func start() {
	var key int

	for {
		utils.CallClear()
		fmt.Println("----------------聊天系统---------------")
		fmt.Println("1) 登录")
		fmt.Println("2) 注册")
		fmt.Println("0) 退出")
		fmt.Printf("请选择(0-2):")

		fmt.Scanln(&key)
		switch key {
		case 1:
			login := &process.UserProcess{}
			login.Login()
			utils.WaitInput()
		case 2:
			//register()
			utils.WaitInput()
		case 0:
			fmt.Println("退出")
			os.Exit(0)
		default:
			fmt.Println("输入有误, 请重新输入!")
			utils.WaitInput()
		}
	}
}
