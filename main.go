package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func VarsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	thing := vars["thing"]
	id := vars["id"]

	fmt.Fprintf(w, "You asked for %s number %s\n", thing, id)
}

func TokenHandler(w http.ResponseWriter, r *http.Request) {
	task := r.URL.Query().Get("task")
	tmpl := template.Must(template.ParseFiles("layout.html"))
	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: task + "1", Done: false},
			{Title: task + "2", Done: true},
			{Title: task + "3", Done: true},
		},
	}
	tmpl.Execute(w, data)
}

func main() {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("static/"))

	r.Handle("/static/", http.StripPrefix("/static/", fs))
	r.HandleFunc("/{thing}/{id}", VarsHandler)
	r.HandleFunc("/", TokenHandler).Host("localhost")

	http.ListenAndServe(":80", r)
}
