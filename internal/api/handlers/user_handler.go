package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/guilhermecosales/security-service/internal/domain/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(r chi.Router, s *service.UserService) {
	h := &UserHandler{service: s}

	r.Route("/users", func(r chi.Router) {
		r.Post("/", h.CreateUser)
		r.Get("/", h.ListUsers)
		r.Get("/{id}", h.GetUserByID)
		r.Put("/{id}", h.UpdateUser)
		r.Patch("/{id}", h.PartialUpdateUser)
		r.Delete("/", h.DeleteUser)
	})
}

// Métodos do handler
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User created"))
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
