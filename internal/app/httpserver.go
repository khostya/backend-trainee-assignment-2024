package app

import (
	"backend-trainee-assignment-2024/config"
	"backend-trainee-assignment-2024/internal/httpserver/handler"
	"backend-trainee-assignment-2024/internal/usecase"
	"backend-trainee-assignment-2024/pkg/httpserver"
	"github.com/go-chi/chi/v5"
)

func newHttpServer(http config.HTTP, dependencies usecase.Dependencies) *httpserver.Server {
	r := chi.NewRouter()

	useCases := usecase.NewUseCases(dependencies)
	handlers.NewRouter(r, useCases)

	return httpserver.New(r,
		httpserver.Port(http.Port),
		httpserver.MaxHeaderBytes(http.MaxHeaderBytes),
		httpserver.IdleTimeout(http.IdleTimeout),
		httpserver.WriteTimeout(http.WriteTimeout),
		httpserver.ReadTimeout(http.ReadTimeout),
	)
}
