package dto

type (
	AuthenticationRequest struct {
		Email    string `json:"email" validate:"required,email,max=150"`
		Password string `json:"password" validate:"required,min=3,max=50"`
	}

	AuthenticationResponse struct {
		AccessToken string `json:"access_token"`
	}
)
