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
