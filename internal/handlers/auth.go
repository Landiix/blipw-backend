package handlers

import (
	"blipw/internal/repository"
	"encoding/json"
	"log"
	"net/http"
	"unicode"
)

type AuthHandler struct {
	repo *repository.AuthRepository
}

func (h *AuthHandler) UserExists(username string) (bool, error) {
	exists := false
	req := "SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)"

	err := h.repo.QueryRowContext(req, username).Scan(&exists)
	if err != nil {
		return false, err
	} else {
		return exists, nil
	}
}

func (h *AuthHandler) hasDigitsAndLetters(password string) bool {
	haveLetter := false
	haveDigit := false

	for _, letter := range password {
		if unicode.IsDigit(letter) {
			haveDigit = true
		} else if unicode.IsLetter(letter) {
			haveLetter = true
		}
		if haveLetter && haveDigit {
			return true
		}
	}
	return false
}

func (h *AuthHandler) register(w http.ResponseWriter, r *http.Request) {

	var data struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil { // 400/1
		log.Printf("Error400/1: %v", err)
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return //возвращаем ошибку 400
	}

	if len(data.Username) < 4 {
		http.Error(w, "Имя должно состоять из 4 и более символов", http.StatusBadRequest)
		return //возвращаем ошибку 400
	}

	exists, err := h.UserExists(data.Username)
	if err != nil { // 500/1
		log.Printf("Error500/1: %v", err)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return //возвращаем ошибку 500
	}
	if exists {
		http.Error(w, "Имя пользователя уже существует", http.StatusConflict)
		return //возвращаем ошибку 409
	}

	if len(data.Password) < 8 {
		http.Error(w, "Пароль должен состоять минимум из 8 символов", http.StatusBadRequest)
		return //возвращаем ошибку 400
	}

	if !h.hasDigitsAndLetters(data.Password) {
		http.Error(w, "В пароле должны быть цифры и символы", http.StatusBadRequest)
		return //возвращаем ошибку 400
	}

	// этапы безопасности пройдены, далее надо будет хешировать пароль, сохранить данные в БД и вернуть ответ 201
}
