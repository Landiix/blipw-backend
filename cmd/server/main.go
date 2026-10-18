package main

import (
	"time"

	"blipw/internal/auth"
	"blipw/internal/config"
	"blipw/internal/database"
	"blipw/internal/handlers"
	"blipw/internal/repository"
	"blipw/internal/service"

	"log"
	"net/http"
)

func main() {
	cfg := config.Load()

	pool, err := database.NewPostgresPool(cfg)
	if err != nil {
		log.Fatalf("Unable to load db: %v", err)
	}
	defer pool.Close()

	log.Println("Succesful connection to db!")

	repo := repository.NewTweetRepository(pool)
	user := repository.NewUserRepository(pool)

	tokens := auth.NewTokenManager(cfg.JWTSecret, 24*time.Hour)
	authService := service.NewAuthService(user, tokens)

	tweetHandler := handlers.NewTweetHandler(repo)
	userHandler := handlers.NewAuthHandler(authService)

	mux := http.NewServeMux()

	protected := auth.RequireAuth(tokens)
	mux.Handle("GET /api/auth/me", protected(http.HandlerFunc(userHandler.Me)))

	mux.HandleFunc("GET /api/tweets", tweetHandler.GetAll)
	mux.HandleFunc("POST /api/tweets", tweetHandler.Create)

	mux.HandleFunc("POST /api/auth/register", userHandler.Register)
	mux.HandleFunc("POST /api/auth/login", userHandler.Login)

	log.Println("Routes registered:")
	log.Println("  GET  /api/tweets -> GetAll")
	log.Println("  POST /api/tweets -> Create")
	log.Println("  POST /api/auth/register -> Register")
	log.Println("  POST /api/auth/login -> Login")

	log.Println("Starting server on :8081...")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	} //http://localhost:8081/api/tweets
}
