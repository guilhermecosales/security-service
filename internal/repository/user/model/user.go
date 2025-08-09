package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID             uuid.UUID `db:"user_id"`
	FirstName          string    `db:"first_name"`
	LastName           string    `db:"last_name"`
	Email              string    `db:"email"`
	Password           string    `db:"password"`
	Locked             bool      `db:"locked"`
	CredentialsExpired bool      `db:"credentials_expired"`
	Enabled            bool      `db:"enabled"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`
}
