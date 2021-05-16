package process

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/gostudy/project2/common/message"
	"github.com/gostudy/project2/common/utils"
	"github.com/gostudy/project2/server/model"
)

type UserProcess struct {
	Conn net.Conn
}

func (this *UserProcess) ServerProcessLogin(msg *message.Message) (err error) {

	var loginMsg message.LoginMsg

	err = json.Unmarshal([]byte(msg.Data), &loginMsg)

	if err != nil {

		return

	}

	var resMsg message.Message

	resMsg.Type = message.LoginResMsgType

	var loginResMsg message.LoginResMsg

	user, err := model.MyUserDao.Login(loginMsg.UserId, loginMsg.UserPwd)
	if err != nil {
		if err == model.ErrorUserNoExitsts {
			loginResMsg.Code = 500
			loginResMsg.Error = err.Error()
		} else if err == model.ErrorUserPwd {
			loginResMsg.Code = 403
			loginResMsg.Error = err.Error()
		} else {
			loginResMsg.Code = 505
			loginResMsg.Error = err.Error()
		}

	} else {
		loginResMsg.Code = 200
		fmt.Println("登录成功", user)
	}
	// 登录处理
	//if loginMsg.UserId == "100" && loginMsg.UserPwd == "123456" {

	//loginResMsg.Code = 200

	//} else {

	//loginResMsg.Code = 500

	//loginResMsg.Error = "该用户不存在, 请注册!"
	//}

	data, err := json.Marshal(loginResMsg)

	if err != nil {

		fmt.Println("json.Marshal fail:", err)

		return
	}

	resMsg.Data = string(data)

	data, err = json.Marshal(resMsg)

	if err != nil {

		fmt.Println("json.Marshal fail:", err)

		return
	}

	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)

	return
}
