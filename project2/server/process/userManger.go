package process

import "errors"

type UserManger struct {
	OnlineUsers map[string]*UserProcess
}

var (
	UuserManger *UserManger
)

func init() {
	UuserManger = &UserManger{
		OnlineUsers: make(map[string]*UserProcess),
	}
}

func (this *UserManger) AddOnlineUser(up *UserProcess) {
	this.OnlineUsers[up.UserId] = up
}

func (this *UserManger) DelOnlineUser(userId string) {
	delete(this.OnlineUsers, userId)
}

func (this *UserManger) GetAllOnlineUser() map[string]*UserProcess {
	return this.OnlineUsers
}

func (this *UserManger) GetOnlineUserById(userId string) (up *UserProcess,
	err error) {

	up, ok := this.OnlineUsers[userId]
	if !ok {
		err = errors.New("不存在该用户, 或者不在线!")
		return
	}
	return
}
