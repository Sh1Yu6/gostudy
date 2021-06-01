package model

import (
	"net"

	"github.com/gostudy/project2/common/message"
)

type CurUser struct {
	Conn net.Conn
	message.User
}
