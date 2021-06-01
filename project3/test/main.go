package main

import (
	"html/template"
	"net/http"
)

func RangeHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html"))
	s := make([]string, 0)
	s = append(s, "asd")
	s = append(s, "klfgj")
	s = append(s, "askldfjk")
	t.Execute(w, s)
}

func WithHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html"))
	t.Execute(w, "heihei")
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html"))
	t.Execute(w, "asd")
}

func main() {

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/range", RangeHandler)
	http.HandleFunc("/with", WithHandler)

	http.ListenAndServe(":8888", nil)
}
