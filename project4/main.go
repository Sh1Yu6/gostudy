package main

import (
	"fmt"
	"net/http"

	"github.com/gostudy/project4/handler"
)

func main() {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	http.HandleFunc("/file/query", handler.FileQueryHandler)
	http.HandleFunc("/file/download", handler.DownloadHandler)
	http.HandleFunc("/file/update", handler.FileMetaUpdateHandler)
	http.HandleFunc("/file/delete", handler.FileDeleteHanler)
	http.HandleFunc("/file/signup", handler.SignupHanlder)

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("Failed to start server, err:", err)
	}
}
