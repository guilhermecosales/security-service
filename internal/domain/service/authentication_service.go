package service

import (
	"context"
	"errors"

	"github.com/guilhermecosales/security-service/internal/domain/model"
	"github.com/guilhermecosales/security-service/internal/infrastructure/repository"
	"github.com/rs/zerolog/log"
)

type AuthenticationService struct {
	userRepository             *repository.UserRepository
	authenticationTokenService *AuthenticationTokenService
}

var (
	ErrInvalidUserCredentials = errors.New("invalid user credentials")
)

func NewAuthenticationService(repository *repository.UserRepository, tokenService *AuthenticationTokenService) *AuthenticationService {
	return &AuthenticationService{
		userRepository:             repository,
		authenticationTokenService: tokenService,
	}
}

func (s *AuthenticationService) GenerateToken(ctx context.Context, userCredentials *model.UserCredentials) (*model.AccessTokenData, error) {
	userDetails := s.loadUserByUsername(ctx, userCredentials.Email)
	if userDetails == nil {
		return nil, ErrInvalidUserCredentials
	}

	accessToken, err := s.authenticationTokenService.GenerateToken(userDetails)
	if err != nil {
		return nil, err
	}

	return &model.AccessTokenData{
		AccessToken: accessToken,
	}, nil
}

func (s *AuthenticationService) loadUserByUsername(ctx context.Context, email string) *model.User {
	userDetails, err := s.userRepository.GetUserByEmail(ctx, email)

	if errors.Is(err, repository.ErrUserNotFound) {
		log.Error().Err(err).Str("email", email).Msg("User not found")
		return nil
	}

	return userDetails
}
