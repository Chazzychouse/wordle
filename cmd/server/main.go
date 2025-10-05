package main

import (
	"net/http"
	"wordle/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	gameHandler := handlers.NewGameHandler()
	r.HandleFunc("/", gameHandler.TestNewWord).Host("localhost")
	http.ListenAndServe(":3000", r)
}
