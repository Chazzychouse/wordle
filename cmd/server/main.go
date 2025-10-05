package main

import (
	"net/http"
	"wordle/internal/handlers"
	"wordle/internal/middleware"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Create handlers
	gameHandler := handlers.NewGameHandler()
	authHandler := handlers.NewAuthHandler()

	// Public routes (no authentication required)
	r.HandleFunc("/auth/url", authHandler.GetAuthURL).Methods("GET")
	r.HandleFunc("/auth/callback", authHandler.HandleCallback).Methods("GET")
	r.HandleFunc("/auth/verify", authHandler.VerifyToken).Methods("GET")

	// Protected routes (authentication required)
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	// TODO:Protect these when OAUTH is implemented
	r.HandleFunc("/game", gameHandler.CreateGame).Methods("POST")
	r.HandleFunc("/game/{id}", gameHandler.GetGame).Methods("GET")
	r.HandleFunc("/games", gameHandler.GetUserGames).Methods("GET")
	r.HandleFunc("/game/{id}/attempt", gameHandler.MakeAttempt).Methods("POST")

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	http.ListenAndServe(":3000", r)
}
