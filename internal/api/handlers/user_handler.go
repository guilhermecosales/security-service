package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/guilhermecosales/security-service/internal/api/dto"
	"github.com/guilhermecosales/security-service/internal/api/helper"
	"github.com/guilhermecosales/security-service/internal/api/mapper"
	"github.com/guilhermecosales/security-service/internal/domain/service"
)

type UserHandler struct {
	service   *service.UserService
	validator *validator.Validate
}

func NewUserHandler(r chi.Router, s *service.UserService) {
	h := &UserHandler{
		service:   s,
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}

	r.Route("/users", func(r chi.Router) {
		r.Post("/", h.CreateUser)
		r.Get("/", h.ListUsers)
		r.Get("/{id}", h.GetUserByID)
		r.Patch("/{id}", h.PartialUpdateUser)
		r.Delete("/{id}", h.DeleteUser)
	})
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var requestUser dto.Request

	if err := json.NewDecoder(r.Body).Decode(&requestUser); err != nil {
		helper.WriteResponse(w, http.StatusBadRequest, nil)
		return
	}

	if err := h.validator.Struct(requestUser); err != nil {
		helper.WriteResponse(w, http.StatusBadRequest, nil)
		return
	}

	userModel := mapper.RequestToModel(&requestUser)

	createdUser, err := h.service.CreateUser(context.Background(), userModel)
	if err != nil {
		helper.WriteResponse(w, http.StatusInternalServerError, nil)
		return
	}

	userResponse := mapper.ModelToResponse(createdUser)
	helper.WriteResponse(w, http.StatusCreated, userResponse)
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List users"))
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get user by ID"))
}

func (h *UserHandler) PartialUpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Partial update user"))
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	pathValue := r.PathValue("id")
	if pathValue == "" {
		helper.WriteResponse(w, http.StatusBadRequest, nil)
		return
	}

	userID, err := uuid.Parse(pathValue)
	if err != nil {
		helper.WriteResponse(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid User Identification",
		})
		return
	}

	err = h.service.DeleteUser(r.Context(), userID)
	if err != nil {
		helper.WriteResponse(w, http.StatusInternalServerError, nil)
		return
	}

	helper.WriteResponse(w, http.StatusNoContent, nil)
	return
}
