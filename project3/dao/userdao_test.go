package dao

import (
	"fmt"
	"testing"
)

func TestUserdao(t *testing.T) {
	fmt.Println("start testing")
	t.Run("验证用户名或者密码: ", testLogin)
	t.Run("验证用户名: ", testRegist)
	//t.Run("保存用户: ", testSave)
}

func testLogin(t *testing.T) {
	user, _ := CheckUserNameAndPassword("admin123", "123456")
	fmt.Println("user info: ", user)
}

func testRegist(t *testing.T) {
	user, _ := CheckUserName("admin123")
	fmt.Println("user info: ", user)
}

func testSave(t *testing.T) {
	SavaUser("admin123", "123456", "admin@qq.com")
}
