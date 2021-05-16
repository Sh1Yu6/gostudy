package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

var serverIP string
var serverPort int

func init() {
	flag.StringVar(&serverIP, "ip", "127.0.0.1", "moren")
	flag.IntVar(&serverPort, "port", 8888, "moren")
}

type Client struct {
	ServerIP   string
	ServerPort int
	Name       string
	conn       net.Conn
	num        int
}

func NewClient(serverIP string, serverPort int) *Client {
	client := &Client{
		ServerIP:   serverIP,
		ServerPort: serverPort,
		num:        999,
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIP, serverPort))
	if err != nil {
		fmt.Println("net dial error: ", err)
		return nil
	}
	client.conn = conn
	return client
}

func main() {
	flag.Parse()

	client := NewClient(serverIP, serverPort)
	if client == nil {
		fmt.Println("连接服务器失败")
		return
	}
	fmt.Println("连接服务器成功")
	go client.DealResponse()
	client.run()
	//select {}
}

func (t *Client) menu() bool {
	fmt.Println("1. 群聊")
	fmt.Println("2. 私聊")
	fmt.Println("3. 更名")
	fmt.Println("0. 退出")
	var selectNum int
	fmt.Scanln(&selectNum)
	if selectNum >= 0 && selectNum <= 3 {
		t.num = selectNum
		return true
	} else {
		fmt.Println("input 0-3")
		return false
	}
}

func (t *Client) run() {
	for t.num != 0 {
		for t.menu() != true {

		}
		switch t.num {
		case 1:
			t.PublicChat()
		case 2:

		case 3:
			t.UpdataName()
		}
	}
	t.conn.Close()
}

func (t *Client) UpdataName() bool {
	fmt.Println("input user name:")
	fmt.Scanln(&t.Name)
	sendMsg := "rename|" + t.Name
	_, err := t.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn write error:", err)
		return false
	}
	return true
}

func (t *Client) PublicChat() {

	var chatMsg string
	for chatMsg != "exit" {
		chatMsg = ""
		fmt.Println("input content(exit out):")
		fmt.Scanln(&chatMsg)
		if len(chatMsg) != 0 {
			sendMsg := chatMsg
			_, err := t.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn write error: ", err)
				break
			}
		}
	}

}

func (t *Client) DealResponse() {
	io.Copy(os.Stdout, t.conn)
}
