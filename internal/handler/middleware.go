package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const (
	userCtxKey = "userID"
)

func (h *Handler) userIdentification(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Empty auth header", http.StatusUnauthorized)
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			http.Error(w, "Invalid auth header", http.StatusUnauthorized)
			return
		}

		if len(headerParts[1]) == 0 {
			http.Error(w, "Token is empty", http.StatusUnauthorized)
			return
		}
		userID, err := h.service.AuthService.ParseToken(headerParts[1])

		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userCtxKey, userID)
		fmt.Println(ctx)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserId(r *http.Request) (int, error) {
	id, ok := r.Context().Value(userCtxKey).(int)
	if !ok {
		return 0, errors.New("User id not found")
	}
	return id, nil
}
