package process

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/gostudy/project2/common/message"
	"github.com/gostudy/project2/common/utils"
)

func ShowMenu() {
	smsProcess := &SmsProcess{}
	for {
		utils.CallClear()
		fmt.Println("----------------聊天系统---------------")
		fmt.Println("1) 显示在线用户列表")
		fmt.Println("2) 发送消息")
		fmt.Println("3) 信息列表")
		fmt.Println("4) 退出系统")
		fmt.Print("请选择(0-3):")

		var key int
		key = 999

		fmt.Scanln(&key)
		switch key {
		case 1:
			outputOnlineUser()
		case 2:
			fmt.Print("想要群发的消息:")
			var content string
			fmt.Scanln(&content)
			smsProcess.SendGroupMsg(content)
		case 3:
		case 0:
			fmt.Println("退出")
			return
		default:
			fmt.Println("输入错误, 请重新输入!")
		}
		utils.WaitInput()
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

		switch msg.Type {
		case message.NotifyUserStatusMsgType:

			var notifyUserStatusMsg message.NotifyUserStatusMsg
			err = json.Unmarshal([]byte(msg.Data), &notifyUserStatusMsg)
			updateUserStatus(&notifyUserStatusMsg)

		case message.SmsMsgType:
			outputGroupMsg(&msg)

		default:
			fmt.Println("服务器返回来未知的消息类型", msg)
		}
	}
}
