package mapper

import (
	"time"

	"github.com/guilhermecosales/security-service/internal/api/dto"
	"github.com/guilhermecosales/security-service/internal/domain/model"
)

func RequestToModel(userRequest *dto.Request) *model.User {
	return &model.User{
		FirstName:          userRequest.FirstName,
		LastName:           userRequest.LastName,
		Email:              userRequest.Email,
		Password:           userRequest.Password,
		Enabled:            userRequest.Enabled,
		Locked:             true,
		CredentialsExpired: true,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}
}

func ModelToResponse(userModel *model.User) *dto.Response {
	return &dto.Response{
		UserID:    userModel.UserID,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
		Email:     userModel.Email,
		Enabled:   userModel.Enabled,
	}
}
