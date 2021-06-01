package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", ff)
	http.ListenAndServe(":8888", nil)
}

func ff(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html"))

	a := []Person{
		{
			"aaa",
			"bbb",
			[]string{"asd", "123", "zxc"},
		},
		{
			"哈哈",
			"暨大",
			[]string{"英雄联盟", "足球", "篮球"},
		},
	}
	t.Execute(w, a)
}

type Person struct {
	Name    string
	School  string
	Hobbies []string
}
