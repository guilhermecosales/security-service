package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/guilhermecosales/security-service/internal/api/dto"
	suite "github.com/guilhermecosales/security-service/internal/api/http"
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
		r.Put("/{id}", h.UpdateUser)
		r.Patch("/{id}", h.PartialUpdateUser)
		r.Delete("/", h.DeleteUser)
	})
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var requestUser dto.Request

	if err := json.NewDecoder(r.Body).Decode(&requestUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if err := h.validator.Struct(requestUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	userModel := mapper.RequestToModel(&requestUser)
	createdUser, err := h.service.CreateUser(context.Background(), userModel)

	if err != nil {
		suite.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	suite.WriteResponse(w, http.StatusOK, createdUser)
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List users"))
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get user by ID"))
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user"))
}

func (h *UserHandler) PartialUpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Partial update user"))
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}
