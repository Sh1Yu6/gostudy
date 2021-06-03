package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/p_test")
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// create a table;
	//_, err = db.Exec(
	//"create table stu( id int primary key,name varchar(10))",
	//)
	//if err != nil {
	//fmt.Println(err)
	//os.Exit(1)
	//}

	//query
	stmt, err := db.Prepare("select name, age from students")
	rows, err := stmt.Query()
	for rows.Next() {
		var name string
		var age int64
		rows.Scan(&name, &age)
		fmt.Println(name, age)
	}

	//other operation insert update delete
	stmt, err := db.Prepare(
		"insert into students (name, age, height) values(?, ?, ?)",
	)
	res, err := stmt.Exec("wang", 34, 1.77)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(res.RowsAffected())
}
