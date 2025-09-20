package mapper

import (
	"github.com/guilhermecosales/security-service/internal/api/dto"
	"github.com/guilhermecosales/security-service/internal/domain/model"
)

func AuthenticationRequestToModel(userRequest *dto.AuthenticationRequest) *model.UserCredentials {
	return &model.UserCredentials{
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}
}

func ModelToAuthenticationResponse(accessToken *model.AccessTokenData) *dto.AuthenticationResponse {
	return &dto.AuthenticationResponse{
		AccessToken: accessToken.AccessToken,
	}
}
