package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/guilhermecosales/security-service/internal/handlers"
)

func HealthRoutes(r chi.Router) {
	r.Route("/health", func(r chi.Router) {
		r.Get("/", handlers.HealthCheck)
	})
}
