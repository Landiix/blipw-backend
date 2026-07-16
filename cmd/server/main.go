package main

import (
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

	auth := service.NewAuthService(user)

	tweetHandler := handlers.NewTweetHandler(repo)
	userHandler := handlers.NewAuthHandler(auth)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/tweets", tweetHandler.GetAll)
	mux.HandleFunc("POST /api/tweets", tweetHandler.Create)

	mux.HandleFunc("POST /api/auth/register", userHandler.Register)

	log.Println("Routes registered:")
	log.Println("  GET  /api/tweets -> GetAll")
	log.Println("  POST /api/tweets -> Create")
	log.Println("  POST /api/auth/register -> Register")

	log.Println("Starting server on :8081...")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	} //http://localhost:8081/api/tweets
}
