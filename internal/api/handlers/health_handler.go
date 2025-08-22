package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/guilhermecosales/security-service/internal/api/helper"
)

type HealthHandler struct {
}

func NewHealthHandler(r chi.Router) {
	r.Route("/health", func(r chi.Router) {
		r.Get("/", healthCheck)
	})
}

func healthCheck(w http.ResponseWriter, _ *http.Request) {
	helper.WriteResponse(w, http.StatusOK, map[string]interface{}{
		"health": "OK",
	})
}
