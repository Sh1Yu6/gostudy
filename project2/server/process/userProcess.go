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
	Conn   net.Conn
	UserId string
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

	_, err = model.MyUserDao.Login(loginMsg.UserId, loginMsg.UserPwd)
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
		this.UserId = loginMsg.UserId
		UuserManger.AddOnlineUser(this)
		this.NotifyOthersOnlineUser(this.UserId)
		for id, _ := range UuserManger.OnlineUsers {
			loginResMsg.UsersId = append(loginResMsg.UsersId, id)
		}
	}

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

func (this *UserProcess) ServerProcessRegister(msg *message.Message) (err error) {

	var registerMsg message.RegisterMsg

	err = json.Unmarshal([]byte(msg.Data), &registerMsg)

	if err != nil {

		return

	}

	var resMsg message.Message

	resMsg.Type = message.RegisterResMsgType

	var registerResMsg message.RegisterResMsg

	err = model.MyUserDao.Register(&registerMsg.User)
	if err != nil {
		if err == model.ErrorUserExitsts {
			registerResMsg.Code = 505
			registerResMsg.Error = err.Error()
		} else {
			registerResMsg.Code = 506
			registerResMsg.Error = err.Error()
		}

	} else {
		registerResMsg.Code = 200

	}

	data, err := json.Marshal(registerResMsg)

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

func (this *UserProcess) NotifyOthersOnlineUser(userId string) {
	for id, up := range UuserManger.OnlineUsers {
		if id == userId {
			continue
		}
		up.notifyMeOnline(userId)
	}
}

func (this *UserProcess) notifyMeOnline(userId string) {
	var msg message.Message
	msg.Type = message.NotifyUserStatusMsgType

	var notifyUserStatusMsg message.NotifyUserStatusMsg
	notifyUserStatusMsg.UserId = userId
	notifyUserStatusMsg.Status = message.UserOnline

	data, err := json.Marshal(notifyUserStatusMsg)
	if err != nil {
		fmt.Println("json.Marshal err", err)
		return
	}

	msg.Data = string(data)

	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal err", err)
		return
	}

	tf := &utils.Transfer{
		Conn: this.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("notifi err", err)
		return
	}
	return

}
