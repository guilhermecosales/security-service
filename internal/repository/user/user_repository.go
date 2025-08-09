package user

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/guilhermecosales/security-service/internal/repository"
	"github.com/guilhermecosales/security-service/internal/repository/user/model"
)

type UserRepository struct {
	db *sql.DB
}

var _ repository.UserRepository = (*UserRepository)(nil)

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user model.User) (*model.User, error) {
	query := `
		INSERT INTO users (
			user_id, first_name, last_name, email, password, locked, credentials_expired, enabled
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING user_id, first_name, last_name, email, password, locked, credentials_expired, enabled
	`

	var createdUser model.User
	err := r.db.QueryRowContext(ctx, query,
		user.UserID,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.Locked,
		user.CredentialsExpired,
		user.Enabled,
	).Scan(
		&createdUser.UserID,
		&createdUser.FirstName,
		&createdUser.LastName,
		&createdUser.Email,
		&createdUser.Password,
		&createdUser.Locked,
		&createdUser.CredentialsExpired,
		&createdUser.Enabled,
	)
	if err != nil {
		return nil, err
	}

	return &createdUser, nil
}

func (r *UserRepository) GetUser(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	return nil, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, userID uuid.UUID, user model.User) (*model.User, error) {
	return nil, nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	return nil
}
