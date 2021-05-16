package process

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"

	"github.com/gostudy/project2/common/message"
	"github.com/gostudy/project2/common/utils"
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

	err := this.LoginVerify(userId, userPwd)
	if err != nil {
		fmt.Println("loginVerify err:", err)
	}
}

func (this *UserProcess) LoginVerify(userId string, userPwd string) (err error) {

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
		go serverProcessMsg(conn)
		ShowMenu()
	} else if loginResMsg.Code == 500 {
		err = errors.New(loginResMsg.Error)
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
