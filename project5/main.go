package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", f)
	http.ListenAndServe(":8888", nil)
}

func f(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(data)
}
