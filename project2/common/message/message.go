package message

const (
	LoginMsgType            = "LoginMsg"
	LoginResMsgType         = "LoginResMsg"
	RegisterMsgType         = "RegisterMsg"
	RegisterResMsgType      = "RegisterResMsg"
	NotifyUserStatusMsgType = "NotifyUserStatusMsg"
	SmsMsgType              = "SmsMsg"
)

const (
	UserOnline     = "zaixian"
	UserOffline    = "buzaixian"
	UserBusyStatus = "mang"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMsg struct {
	UserId   string `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMsg struct {
	Code    int      `json:"code"`
	UsersId []string `json:"usersId"`
	Error   string   `json:"error"`
}

type RegisterMsg struct {
	User User `json:"user"`
}

type RegisterResMsg struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type NotifyUserStatusMsg struct {
	UserId string `json:"userId"`
	Status string `json:"status"`
}

type SmsMsg struct {
	Content string `json:"content"`
	User
}

type SmsResMsg struct {
}
