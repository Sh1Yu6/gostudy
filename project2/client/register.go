package process

import "fmt"

func register() {
	loopRegister := true
	for loopRegister {
		callClear()
		fmt.Println("----------------聊天系统---------------")
		fmt.Println("------------------注册-----------------")
		fmt.Printf("请输入账号:")
		var userId string
		fmt.Scanln(&userId)

		var userPwd string
		var userPwd2 string
		loopPwd := true
		for loopPwd {
			fmt.Printf("请输入密码:")
			fmt.Scanln(&userPwd)

			fmt.Printf("请确认密码:")
			fmt.Scanln(&userPwd2)

			if userPwd == userPwd2 {
				loopPwd = false
			} else {
				fmt.Println("两次密码不一样, 请重新输入!")
			}
		}

		err := loginVerify(userId, userPwd)
		if err != nil {
			fmt.Println("账户已存在!")
			waitInput()
			continue
		}
		break
	}
}

func registerVerify(userId string, userPwd string) (err error) {
	fmt.Println("ok")
	return nil
}
