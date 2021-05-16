package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"

	"github.com/gostudy/project2/common/message"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (this *Transfer) ReadPkg() (msg message.Message, err error) {

	// 先接收包头, 说明接下来的数据体量有多少
	_, err = this.Conn.Read(this.Buf[:4])

	if err != nil {
		return
	}

	var pkgLen uint32

	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])

	// 真正接收数据
	n, err := this.Conn.Read(this.Buf[:pkgLen])

	if n != int(pkgLen) || err != nil {

		err = errors.New("read data error")

		return
	}

	err = json.Unmarshal(this.Buf[:pkgLen], &msg)

	if err != nil {

		fmt.Println("json.Unmarshal err:", err)

		return
	}

	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {

	// 协定传送的数据有多大
	var pkgLen uint32

	pkgLen = uint32(len(data))

	binary.BigEndian.PutUint32(this.Buf[:4], pkgLen)

	n, err := this.Conn.Write(this.Buf[:4])

	if n != 4 || err != nil {

		fmt.Println("conn.Write(bytes) err:", err)

		return
	}

	// 真正传输数据
	n, err = this.Conn.Write(data)

	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.wiRite(data) fail:", err)
		return
	}

	return
}

var Clear map[string]func()

func init() {
	Clear = make(map[string]func())
	Clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	Clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := Clear[runtime.GOOS]
	if ok {
		value()
	} else {
		fmt.Println("不能应对你的操作系统!")
		os.Exit(1)
	}

}

func WaitInput() {
	fmt.Printf("请按任意键继续...")
	fmt.Scanln()
}
