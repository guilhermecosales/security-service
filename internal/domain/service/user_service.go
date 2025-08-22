package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/guilhermecosales/security-service/internal/domain/model"
	"github.com/guilhermecosales/security-service/internal/infrastructure/repository"
	"github.com/rs/zerolog/log"
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(repository *repository.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	return s.repository.CreateUser(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	if err := s.repository.DeleteUser(ctx, userID); err != nil {
		log.Err(err).Str("userId", userID.String()).Msg("failed to delete user")
		return err
	}

	return nil
}
