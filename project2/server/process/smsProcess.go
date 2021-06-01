package process

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/gostudy/project2/common/message"
	"github.com/gostudy/project2/common/utils"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMsg(msg *message.Message) (err error) {

	var smsMsg message.SmsMsg
	err = json.Unmarshal([]byte(msg.Data), &smsMsg)
	if err != nil {
		fmt.Println("json marshal err:", err)
		return
	}

	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("json marshal err:", err)
		return
	}

	for id, uup := range UuserManger.OnlineUsers {
		if id == smsMsg.UserId {
			continue
		}
		this.SendMsgToEachOnlineUser(data, uup.Conn)
	}

	return
}

func (this *SmsProcess) SendMsgToEachOnlineUser(data []byte, conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败 err", err)
	}
}
