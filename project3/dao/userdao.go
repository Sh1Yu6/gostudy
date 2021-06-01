package dao

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gostudy/project3/model"
	"github.com/gostudy/project3/utils"
)

func CheckUserNameAndPassword(username, password string) (
	*model.User, error) {

	sqlStr := "select id, username, password, email from users where username = ? and password = ?"
	row := utils.Db.QueryRow(sqlStr, username, password)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, nil
}

func CheckUserName(username string) (
	*model.User, error) {

	sqlStr := "select id, username, password, email from users"
	sqlStr += " where username = ?"
	row := utils.Db.QueryRow(sqlStr, username)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, nil
}

func SavaUser(username, password, email string) error {
	sqlStr := "insert into users(username, password, email) values(?,?,?)"

	_, err := utils.Db.Exec(sqlStr, username, password, email)
	if err != nil {
		fmt.Println("aaaaaaaaa", err)
		return err
	}
	return nil
}
