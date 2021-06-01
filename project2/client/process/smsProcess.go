package process

import (
	"encoding/json"
	"fmt"

	"github.com/gostudy/project2/common/message"
	"github.com/gostudy/project2/common/utils"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMsg(content string) (err error) {
	var msg message.Message
	msg.Type = message.SmsMsgType

	var smsMsg message.SmsMsg
	smsMsg.Content = content
	smsMsg.UserId = CurUser.UserId
	smsMsg.UserStatus = CurUser.UserStatus

	data, err := json.Marshal(smsMsg)
	if err != nil {
		fmt.Println("SendGroupMsg json marshal err", err)
		return
	}

	msg.Data = string(data)

	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("SendGroupMsg json marshal err", err)
		return
	}

	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMsg writepkg err", err)
		return
	}
	return

}
