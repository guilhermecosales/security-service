package dto

type (
	Request struct {
		FirstName string `json:"first_name" validate:"required,min=1,max=50"`
		LastName  string `json:"last_name" validate:"required,min=1,max=50"`
		Email     string `json:"email" validate:"required,email"`
		Password  string `json:"password" validate:"required"`
		Enabled   bool   `json:"enabled" validate:"required"`
	}
)
