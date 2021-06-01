package db

import (
	"database/sql"
	"fmt"

	"github.com/gostudy/project4/db/mysql"
)

func OnFileUploadFinished(filehash, filename, fileaddr string,
	filesize int64) bool {

	stmt, err := mysql.DBConn().Prepare(
		"insert into tbl_file (file_sha1, file_name, file_size, " +
			"file_addr, status) values(?, ?, ?, ?, 1)")

	if err != nil {
		fmt.Println("Failed to prepare statement, err: ", err)
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(filehash, filename, filesize, fileaddr)

	if err != nil {
		fmt.Println(err)
		return false
	}

	rf, err := ret.RowsAffected()
	if err == nil {
		if rf == 0 {
			fmt.Println("file with hash:", filehash, "has been uploaded before")
		}
	}
	return true
}

type TableFile struct {
	FileHash string
	FileName sql.NullString
	FileSize sql.NullInt64
	FileAddr sql.NullString
}

func GetFileMeta(filehash string) (*TableFile, error) {
	stmt, err := mysql.DBConn().Prepare(
		"select file_sha1, file_addr, file_name, file_size from tbl_file" +
			" where file_sha1 = ? and status = 1 limit 1")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer stmt.Close()

	tfile := &TableFile{}
	err = stmt.QueryRow(filehash).Scan(&tfile.FileHash, &tfile.FileAddr,
		&tfile.FileName, &tfile.FileSize)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return tfile, nil
}
