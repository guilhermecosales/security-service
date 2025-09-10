package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/guilhermecosales/security-service/pkg/protocol"
)

type HealthHandler struct {
}

func NewHealthHandler(r chi.Router) {
	r.Route("/health", func(r chi.Router) {
		r.Get("/", healthCheck)
	})
}

func healthCheck(w http.ResponseWriter, _ *http.Request) {
	protocol.WriteResponse(w, http.StatusOK, map[string]interface{}{
		"health": "OK",
	})
}
