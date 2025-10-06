package main

import (
	"net/http"
	"wordle/internal/handlers"
	"wordle/internal/middleware"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Use(middleware.CORSMiddleware)
	gameHandler := handlers.NewGameHandler()
	authHandler := handlers.NewAuthHandler()

	r.HandleFunc("/auth/url", authHandler.GetAuthURL).Methods("GET")
	r.HandleFunc("/auth/callback", authHandler.HandleCallback).Methods("GET")
	r.HandleFunc("/auth/verify", authHandler.VerifyToken).Methods("GET")

	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	// TODO:Protect these when OAUTH is implemented
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	}).Methods("GET")
	r.HandleFunc("/game", gameHandler.CreateGame).Methods("POST")
	r.HandleFunc("/game/{id}", gameHandler.GetGame).Methods("GET")
	// TODO: Use email from auth middleware
	r.HandleFunc("/games/{email}", gameHandler.GetUserGames).Methods("GET")
	r.HandleFunc("/game/{id}/attempt", gameHandler.MakeAttempt).Methods("POST")

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	http.ListenAndServe(":3333", r)
}
