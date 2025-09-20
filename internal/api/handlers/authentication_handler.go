package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/guilhermecosales/security-service/internal/api/dto"
	"github.com/guilhermecosales/security-service/internal/api/mapper"
	"github.com/guilhermecosales/security-service/internal/domain/service"
	"github.com/guilhermecosales/security-service/pkg/protocol"
)

type AuthenticationHandler struct {
	service   *service.AuthenticationService
	validator *validator.Validate
}

func NewAuthenticationHandler(r chi.Router, s *service.AuthenticationService) {
	h := &AuthenticationHandler{
		service:   s,
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}

	r.Route("/authentication", func(r chi.Router) {
		r.Post("/token", h.generateToken)
	})
}

func (h *AuthenticationHandler) generateToken(w http.ResponseWriter, r *http.Request) {
	var authRequest dto.AuthenticationRequest

	if err := json.NewDecoder(r.Body).Decode(&authRequest); err != nil {
		protocol.WriteResponse(w, http.StatusBadRequest, nil)
		return
	}

	if err := h.validator.Struct(authRequest); err != nil {
		protocol.WriteResponse(w, http.StatusBadRequest, nil)
		return
	}

	userCredentials := mapper.AuthenticationRequestToModel(&authRequest)

	accessTokenData, err := h.service.GenerateToken(r.Context(), userCredentials)
	if err != nil {
		if errors.Is(err, service.ErrInvalidUserCredentials) {
			protocol.WriteResponse(w, http.StatusUnauthorized, map[string]string{"error": err.Error()})
			return
		}

		protocol.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		return
	}

	accessTokenDataResponse := mapper.ModelToAuthenticationResponse(accessTokenData)
	protocol.WriteResponse(w, http.StatusOK, accessTokenDataResponse)
}
