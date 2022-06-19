package main

import (
	"net/http"
	"html/template"
	"log"
)

var tmpl *template.Template

type Todo struct {
	Item string
	Done bool
}

type PageData struct {
	Title string
	Todos []Todo
}



func todo(w http.ResponseWriter, r *http.Request) {
	data := PageData {
		Title: "TODO LIST",
		Todos: []Todo{
			{Item: "Install Go", Done: true},
			{Item: "Learn Go", Done: true},
			{Item: "Make Todo List in Go", Done: false},
		},
	}

	tmpl.Execute(w, data)
}


func main() {
	mux := http.NewServeMux()

	tmpl  = template.Must(template.ParseFiles("templates/index.gohtml"))

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", todo)

	log.Fatal(http.ListenAndServe("0.0.0.0:9090", mux))
}
