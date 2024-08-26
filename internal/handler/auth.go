package handler

import (
	"encoding/json"
	"github.com/fanfaronDo/code_education_api/internal/domain"
	"net/http"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var signUpUser domain.User
	err := json.NewDecoder(r.Body).Decode(&signUpUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.service.AuthService.CreateUser(signUpUser)
	if err != nil {
		http.Error(w, "Could not create user", http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]interface{}{
		"id": id,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Could not send response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var signInUser = struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&signInUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.service.AuthService.GenerateToken(signInUser.Username, signInUser.Password)
	if err != nil {
		http.Error(w, "Could not authorisation", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"token": token,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Could not send response", http.StatusInternalServerError)
		return
	}
}
