package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"unicode/utf8"

	"blipw/internal/repository"
)

type TweetHandler struct {
	repo *repository.TweetRepository
}

type CreateTweetRequest struct {
	UserID  int64  `json:"user_id"`
	Content string `json:"content"`
}

func NewTweetHandler(repo *repository.TweetRepository) *TweetHandler {
	return &TweetHandler{
		repo: repo,
	}
}

func (h *TweetHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Content-Type", "application/json")

	tweets, err := h.repo.GetAll(r.Context())
	if err != nil {
		log.Println("Unable to get data: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tweets)
	log.Printf("tweets from db: %+v\n", tweets)
}

func (h *TweetHandler) Create(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAll called! Method:", r.Method, "URL:", r.URL.Path)
	var req CreateTweetRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if len(req.Content) == 0 {
		http.Error(w, "Length of Content is 0!", http.StatusBadRequest)
		return
	}
	if utf8.RuneCountInString(req.Content) > 280 {
		http.Error(w, "Length of Content > 280 symbols!", http.StatusBadRequest)
		return
	}

	tweet, err := h.repo.Create(r.Context(), req.UserID, req.Content)
	if err != nil {
		log.Println("Unable to create tweet: ", err)
		http.Error(w, "Internal server error!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tweet)
}
