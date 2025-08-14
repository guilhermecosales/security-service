package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/guilhermecosales/security-service/internal/domain/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user model.User) (*model.User, error)
	GetUser(ctx context.Context, userID uuid.UUID) (*model.User, error)
	UpdateUser(ctx context.Context, userID uuid.UUID, user model.User) (*model.User, error)
	DeleteUser(ctx context.Context, userID uuid.UUID) error
}

type RoleRepository interface {
}
