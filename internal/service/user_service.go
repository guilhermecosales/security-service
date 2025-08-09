package service

import (
	"context"

	"github.com/guilhermecosales/security-service/internal/repository/user"
	"github.com/guilhermecosales/security-service/internal/repository/user/model"
)

type UserService struct {
	repository user.UserRepository
}

func NewUserService(repository user.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) CreateUser(ctx context.Context, user model.User) (*model.User, error) {
	return s.repository.CreateUser(ctx, user)
}
