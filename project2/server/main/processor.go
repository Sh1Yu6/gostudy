package main

import (
	"fmt"
	"io"
	"net"

	"github.com/gostudy/project2/common/message"
	"github.com/gostudy/project2/common/utils"
	"github.com/gostudy/project2/server/process"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) serverProcessMsg(msg *message.Message) (err error) {

	switch msg.Type {

	case message.LoginMsgType:

		up := &process.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(msg)

	case message.RegisterMsgType:

		up := &process.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(msg)

	case message.SmsMsgType:
		smsProcess := &process.SmsProcess{}
		smsProcess.SendGroupMsg(msg)

	default:

		fmt.Println("消息类型不存在, 无法处理!")

	}
	return
}

func (this *Processor) process2() error {
	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}

		msg, err := tf.ReadPkg()

		if err != nil {
			if err == io.EOF {

				fmt.Println("正常退出")

			} else {

				fmt.Println("readPkg err", err)

			}

			return err
		}

		err = this.serverProcessMsg(&msg)
		if err != nil {
			return err
		}
	}
}
