package handler

import (
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	// get services
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *chi.Mux {
	route := chi.NewRouter()
	route.Route("/auth", func(c chi.Router) {
		route.Get("/s", h.signIn)
		route.Post("/sign_up", h.signUp)
		route.Post("/sign_in", h.signIn)
	})
	return route
}
