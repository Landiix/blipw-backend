package main

import (
	"blipw/internal/config"
	"blipw/internal/database"
	"blipw/internal/repository"
	"context"

	"log"
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
	tweets, err := repo.GetAll(context.Background())
	if err != nil {
		log.Println("Unable to get data: ", err)
	}
	log.Printf("tweets from db: %+v\n", tweets)
}
