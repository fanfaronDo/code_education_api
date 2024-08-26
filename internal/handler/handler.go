package handler

import (
	"github.com/fanfaronDo/code_education_api/internal/service"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *service.Service
	// get services
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *chi.Mux {
	route := chi.NewRouter()

	route.Route("/auth", func(r chi.Router) {
		r.Post("/sign_up", h.signUp)
		r.Post("/sign_in", h.signIn)
	})

	route.Route("/api", func(r chi.Router) {
		r.Use(h.userIdentification)
		r.Route("/notes", func(notes chi.Router) {
			notes.Post("/", h.createNote)
			notes.Get("/", h.getNotes)
		})

	})

	return route
}
