package main

import (
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, "")
}

func main() {

	http.HandleFunc("/main", IndexHandler)

	http.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("views/static/"))))

	http.Handle("/pages/", http.StripPrefix("/pages/",
		http.FileServer(http.Dir("views/pages/"))))

	http.ListenAndServe(":8888", nil)
}
