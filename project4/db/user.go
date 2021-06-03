package db

import (
	"fmt"

	"github.com/gostudy/project4/db/mysql"
)

func UserSignup(username, passwd string) bool {
	stmt, err := mysql.DBConn().Prepare(
		"insert ignore into tbl_user (user_name, user_pwd) values(?, ?)")
	if err != nil {
		fmt.Println("Falied to insert, err", err)
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(username, passwd)
	if err != nil {
		fmt.Println("Falied to insert, err", err)
		return false
	}

	rowsAffected, err := ret.RowsAffected()
	if err == nil && rowsAffected > 0 {
		return true
	}

	return false
}

func UserSignin(username, encpwd string) bool {
	stmt, err := mysql.DBConn().Prepare(
		"select * from tbl_user where user_name = ? limit 1",
	)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer stmt.Close()

	rows, err := stmt.Query(username)
	if err != nil {
		fmt.Println(err)
		return false
	} else if rows == nil {
		fmt.Println("username not found:", username)
		return false
	}

	pRows := mysql.ParseRows(rows)
	if len(pRows) > 0 && string(pRows[0]["user_pwd"].([]byte)) == encpwd {
		return true
	}
	return false
}

func UpdateToken(username, token string) bool {
	stmt, err := mysql.DBConn().Prepare(
		"replace into tbl_user_token(`user_name`, `user_token`) values(?,?)",
	)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, token)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

type User struct {
	Username     string
	Email        string
	Phone        string
	SignupAt     string
	LastActiveAt string
	Status       int
}

func GetUserInfo(username string) (User, error) {
	user := User{}

	stmt, err := mysql.DBConn().Prepare(
		"select user_name, signup_at from tbl_user where user_name = ? limit 1",
	)
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&user.Username, &user.SignupAt)
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	return user, nil
}
