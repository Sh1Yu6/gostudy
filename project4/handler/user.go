package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gostudy/project4/db"
	"github.com/gostudy/project4/util"
)

const (
	pwdSalt = "*#890"
)

func SignupHanlder(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}

	r.ParseForm()

	username := r.Form.Get("username")
	passwd := r.Form.Get("password")

	if len(username) < 3 || len(passwd) < 5 {
		w.Write([]byte("invalid parameter"))
		return
	}

	encPasswd := util.Sha1([]byte(passwd + pwdSalt))
	suc := db.UserSignup(username, encPasswd)
	if suc {
		w.Write([]byte("SUCCESS"))
	} else {
		fmt.Println("failed")
		w.Write([]byte("FAILED"))
	}
}
