package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	IP        string
	Port      int
	OnlineMap map[string]*User
	mapLock   sync.RWMutex
	Message   chan string
}

//asdfasdf
func NewServer(ip string, port int) *Server {
	return &Server{
		IP:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
}

//qqqqqqqqqqq
func (t *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", t.IP, t.Port))
	if err != nil {
		fmt.Println("net.Listen err: ", err)
	}
	defer listener.Close()

	go t.ListenMessage()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accpet err: ", err)
			continue
		}

		go t.Handler(conn)
	}
}

func (t *Server) Handler(conn net.Conn) {

	user := NewUser(conn, t)

	user.Online()

	isLive := make(chan bool)

	go func() {
		for {
			buf := make([]byte, 4096)
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("conn read err: ", err)
				return
			}
			msg := string(buf[:n])
			user.DoMessage(msg)
			isLive <- true
		}
	}()

	for {
		select {
		case <-isLive:

		case <-time.After(time.Second * 100):
			user.SendMsg("kick out")
			close(user.C)
			conn.Close()
			return
		}

	}
}

func (t *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	t.Message <- sendMsg
}

func (t *Server) ListenMessage() {
	for {
		msg := <-t.Message
		t.mapLock.Lock()
		for _, cli := range t.OnlineMap {
			cli.C <- msg
		}
		t.mapLock.Unlock()
	}
}
