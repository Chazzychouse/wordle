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
	tmpl := template.Must(template.ParseFiles("web/templates/layout.html"))

	if task == "" {
		task = "Do thing"
	}

	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: task + " 1", Done: false},
			{Title: task + " 2", Done: true},
			{Title: task + " 3", Done: true},
		},
	}

	tmpl.Execute(w, data)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", TokenHandler).Host("localhost")
	r.HandleFunc("/give/{thing}/{id}", VarsHandler)

	http.ListenAndServe(":80", r)
}
