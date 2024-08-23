package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Handler struct {
	// get services
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *chi.Mux {
	route := chi.NewRouter()
	route.Get("/", func(w http.ResponseWriter, r *http.Request) {
		response := struct {
			Message string `json:"message"`
		}{
			Message: "Hello World",
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	return route
}
