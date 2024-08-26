package handler

import (
	"github.com/fanfaronDo/code_education_api/internal/service"
	"github.com/go-chi/chi/v5"
	"net/http"
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
	route.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	route.Route("/auth", func(r chi.Router) {
		r.Post("/sign_up", h.signUp)
		r.Post("/sign_in", h.signIn)
	})

	return route
}
