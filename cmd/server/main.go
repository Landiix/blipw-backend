package main

import (
	"blipw/internal/config"
	"blipw/internal/database"
	"blipw/internal/repository"

	"context"
	"encoding/json"
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

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/tweets", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Content-Type", "application/json")

		tweets, err := repo.GetAll(context.Background())
		if err != nil {
			log.Println("Unable to get data: ", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(tweets)
		log.Printf("tweets from db: %+v\n", tweets)
	})

	log.Println("Starting server on :8080...")
	http.ListenAndServe(":8080", mux)
}
