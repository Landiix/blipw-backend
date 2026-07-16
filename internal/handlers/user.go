package handlers

import (
	"encoding/json"
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
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	user, err := h.authService.Register(r.Context(), req.Username, req.Password)
	if err != nil {
		log.Println("Register error:", err)

		if err == service.ErrUserAlreadyExists {
			http.Error(w, "Username already exists", http.StatusConflict) // 409
			return
		}

		http.Error(w, "Internal server error", http.StatusInternalServerError) // 500
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	json.NewEncoder(w).Encode(user)
}
