package handlers

import (
	"backend-trainee-assignment-2024/internal/httpserver/handler/banner"
	"backend-trainee-assignment-2024/internal/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
)

func NewRouter(h *chi.Mux, useCases usecase.UseCases) {
	h.Use(cors.AllowAll().Handler)
	h.Use(middleware.Recoverer)
	h.Use(middleware.RequestID)

	h.Get("/healthz", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	})

	h.Group(func(r chi.Router) {
		banner.New(r, useCases.Banner)
	})
}
