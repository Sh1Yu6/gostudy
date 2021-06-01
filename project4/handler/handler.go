package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gostudy/project4/meta"
	"github.com/gostudy/project4/util"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		data, err := ioutil.ReadFile("./static/view/index.html")

		if err != nil {

			io.WriteString(w, "internal server err")

			return
		}

		io.WriteString(w, string(data))

	} else if r.Method == "POST" {

		r.FormFile("file")

		file, head, err := r.FormFile("file")

		if err != nil {

			fmt.Println("Failed to get data, err: ", err)

			return
		}

		defer file.Close()

		newFile, err := os.Create("./tmp/" + head.Filename)

		if err != nil {
			fmt.Println("Failed to crete file, err: ", err)
			return
		}

		defer newFile.Close()

		size, err := io.Copy(newFile, file)

		if err != nil {
			fmt.Println("Falied to save data into file, err: ", err)
			return
		}

		newFile.Seek(0, os.SEEK_SET)
		fileMeta := meta.FileMeta{
			FileName: head.Filename,
			Location: "./tmp/" + head.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),
			FileSize: size,
			FileSha1: util.FileSha1(newFile),
		}
		//meta.UpdateFileMeta(fileMeta)
		meta.UpdateFileMetaDB(fileMeta)

		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload finished!")
}

func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	filehash := r.Form["filehash"][0]

	fMeta, err := meta.GetFileMetaDB(filehash)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(fMeta)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

func FileQueryHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	limitCnt, _ := strconv.Atoi(r.Form.Get("limit"))

	fileMetas := meta.GetLastFileMetas(limitCnt)
	data, err := json.Marshal(fileMetas)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	filehash := r.Form.Get("filehash")

	fMeta := meta.GetFileMeta(filehash)

	fd, err := os.Open(fMeta.Location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer fd.Close()

	data, err := ioutil.ReadAll(fd)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octect-stream")
	w.Header().Set("Content-disposition", "attachment;filename=\""+fMeta.FileName+"\"")
	w.Write(data)
}

func FileMetaUpdateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	opType := r.Form.Get("op")
	filehash := r.Form.Get("filehash")
	newFileName := r.Form.Get("filename")

	if opType != "0" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	//if r.Method != "POST" {
	//w.WriteHeader(http.StatusMethodNotAllowed)
	//return
	//}

	curFileMeta := meta.GetFileMeta(filehash)

	err := os.Rename(curFileMeta.Location, "./tmp/"+newFileName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	curFileMeta.FileName = newFileName

	meta.UpdateFileMeta(curFileMeta)

	data, err := json.Marshal(curFileMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func FileDeleteHanler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	filehash := r.Form.Get("filehash")

	fMeta := meta.GetFileMeta(filehash)

	err := os.Remove(fMeta.Location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	meta.RemoveFileMeta(filehash)

	w.WriteHeader(http.StatusOK)
}
