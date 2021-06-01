package process

import (
	"encoding/json"
	"fmt"

	"github.com/gostudy/project2/common/message"
)

func outputGroupMsg(msg *message.Message) {

	var smsMsg message.SmsMsg
	err := json.Unmarshal([]byte(msg.Data), &smsMsg)
	if err != nil {
		fmt.Println("json unmarshal err", err)
		return
	}

	info := fmt.Sprintf("%s 说: %s", smsMsg.UserId, smsMsg.Content)
	fmt.Println(info)
}
