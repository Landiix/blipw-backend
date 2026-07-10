package main

import (
	"blipw/internal/config"
	"blipw/internal/database"
	"blipw/internal/handlers"
	"blipw/internal/repository"

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

	tweetHandler := handlers.NewTweetHandler(repo)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/tweets", tweetHandler.GetAll)
	mux.HandleFunc("POST /api/tweets", tweetHandler.Create)

	log.Println("Routes registered:")
	log.Println("  GET  /api/tweets -> GetAll")
	log.Println("  POST /api/tweets -> Create")

	log.Println("Starting server on :8080...")
	http.ListenAndServe(":8080", mux)
}
