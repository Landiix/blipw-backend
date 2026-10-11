package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"blipw/internal/service"
)

type AuthHandler struct {
	authService *service.AuthService
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", 400)
		return
	}

	user, err := h.authService.Register(r.Context(), req.Username, req.Password)
	if err != nil {
		log.Println("Register error:", err)

		if err == service.ErrUserAlreadyExists {
			http.Error(w, "Username already exists", 409)
			return
		}

		http.Error(w, "Internal server error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(user)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", 400)
		return
	}

	user, err := h.authService.Login(r.Context(), req.Username, req.Password)

	if errors.Is(err, service.ErrInvalidCredentials) {
		http.Error(w, "Invalid username or password", 401)
		return
	}

	w.Header().Set("Content type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(user)
}
