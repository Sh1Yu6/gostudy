package process

import (
	"fmt"

	"github.com/gostudy/project2/client/model"
	"github.com/gostudy/project2/common/message"
)

var onlineUsers map[string]*message.User = make(map[string]*message.User)

var CurUser model.CurUser

func updateUserStatus(notifyUserStatusMsg *message.NotifyUserStatusMsg) {

	user, ok := onlineUsers[notifyUserStatusMsg.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMsg.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMsg.Status
	onlineUsers[notifyUserStatusMsg.UserId] = user

	outputOnlineUser()
}

func outputOnlineUser() {
	fmt.Println("\n当前在线用户")
	for id, _ := range onlineUsers {
		fmt.Println("用户: ", id)
	}
}
