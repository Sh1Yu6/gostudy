package main

import (
	"net"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	go user.ListenMessage()
	return &user
}

func (t *User) ListenMessage() {
	for {
		msg := <-t.C
		t.conn.Write([]byte(msg + "\n"))
	}
}

func (t *User) Online() {
	t.server.mapLock.Lock()
	t.server.OnlineMap[t.Name] = t
	t.server.mapLock.Unlock()
	t.server.BroadCast(t, "shangxian")
}

func (t *User) Offline() {
	t.server.mapLock.Lock()
	delete(t.server.OnlineMap, t.Name)
	t.server.mapLock.Unlock()
	t.server.BroadCast(t, "xiaxian")

}

func (t *User) DoMessage(msg string) {
	if msg == "who" {
		t.server.mapLock.Lock()
		for _, user := range t.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "zaixian..\n"
			t.SendMsg(onlineMsg)
		}
		t.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		newName := strings.Split(msg, "|")[1]
		_, ok := t.server.OnlineMap[newName]
		if ok {
			t.SendMsg("当前用户名已经被使用\n")
		} else {
			t.server.mapLock.Lock()
			delete(t.server.OnlineMap, t.Name)
			t.server.OnlineMap[newName] = t
			t.server.mapLock.Unlock()
			t.Name = newName
			t.SendMsg("修改用户名为: " + t.Name + "\n")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			t.SendMsg("aaaaaaaaaaa\n")
		}
		remoteUser, ok := t.server.OnlineMap[remoteName]
		if !ok {
			t.SendMsg("bbbbbbbbbbbb\n")
			return
		}

		content := strings.Split(msg, "|")[2]
		remoteUser.SendMsg(t.Name + "to you: " + content)

	} else {
		t.server.BroadCast(t, msg)
	}
}

func (t *User) SendMsg(msg string) {
	t.conn.Write([]byte(msg))
}
