package process

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/gostudy/project2/common/message"
	"github.com/gostudy/project2/common/utils"
	"github.com/pkg/errors"
)

type UserProcess struct {
}

func (this *UserProcess) Login() {
	utils.CallClear()
	fmt.Println("----------------聊天系统---------------")
	fmt.Println("------------------登录-----------------")

	fmt.Printf("请输入账号:")
	var userId string
	fmt.Scanln(&userId)

	fmt.Printf("请输入密码:")
	var userPwd string
	fmt.Scanln(&userPwd)

	err := this.loginVerify(userId, userPwd)
	if err != nil {
		fmt.Println("loginVerify err:", err)
	}
}

func (this *UserProcess) loginVerify(userId string, userPwd string) (err error) {

	conn, err := net.Dial("tcp", "139.9.5.199:8888")

	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()

	data, err := toJson(userId, userPwd)
	if err != nil {
		return
	}

	tf := utils.Transfer{
		Conn: conn,
	}

	tf.WritePkg(data)

	msg, err := tf.ReadPkg()

	if err != nil {
		return
	}

	var loginResMsg message.LoginResMsg

	err = json.Unmarshal([]byte(msg.Data), &loginResMsg)
	if loginResMsg.Code == 200 {
		fmt.Println("登录成功")

		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline

		for _, v := range loginResMsg.UsersId {
			if v == userId {
				continue
			}
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		utils.WaitInput()
		go serverProcessMsg(conn)
		ShowMenu()
	} else {
		err = errors.New(loginResMsg.Error)
	}

	return
}

func (this *UserProcess) Register() {
	loopRegister := true
	for loopRegister {
		utils.CallClear()
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

		var userName string
		fmt.Printf("请输入用户名:")
		fmt.Scanln(&userName)

		err := this.registerVerify(userId, userPwd, userName)
		if err != nil {
			fmt.Println("账户已存在!")
		}
		break
	}
}

func (this *UserProcess) registerVerify(userId, userPwd,
	userName string) (err error) {

	conn, err := net.Dial("tcp", "139.9.5.199:8888")

	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()

	// 准备要发送的注册数据
	data, err := toJson2(userId, userPwd, userName)
	if err != nil {
		return
	}

	tf := utils.Transfer{
		Conn: conn,
	}

	tf.WritePkg(data)

	msg, err := tf.ReadPkg()

	if err != nil {
		return
	}

	var registerResMsg message.RegisterResMsg

	err = json.Unmarshal([]byte(msg.Data), &registerResMsg)
	if registerResMsg.Code == 200 {
		fmt.Println("注册成功")
	} else {
		err = errors.New(registerResMsg.Error)
		return
	}

	return
}

func toJson(userId, userPwd string) (data []byte, err error) {
	// 数据组织
	var msg message.Message
	msg.Type = message.LoginMsgType

	var loginMsg message.LoginMsg
	loginMsg.UserId = userId
	loginMsg.UserPwd = userPwd

	// 把数据转成json
	data, err = json.Marshal(loginMsg)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}

	msg.Data = string(data)

	// 整个包转成json
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}
	return
}

func toJson2(userId, userPwd, userName string) (data []byte, err error) {
	var msg message.Message
	msg.Type = message.RegisterMsgType

	var registerMsg message.RegisterMsg
	registerMsg.User.UserId = userId
	registerMsg.User.UserPwd = userPwd
	registerMsg.User.UserName = userName

	data, err = json.Marshal(registerMsg)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}

	msg.Data = string(data)

	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}
	return
}
