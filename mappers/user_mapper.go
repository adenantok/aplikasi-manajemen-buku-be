package mappers

import (
	"articlehub-be/dto"
	"articlehub-be/models"
)

func ToUser(userDTO dto.UserDTO) models.User {
	return models.User{
		Username: userDTO.Username,
		Password: userDTO.Password,
		Email:    userDTO.Email,
	}
}

func ToUserLogin(userDTO dto.UserDTOLogin) models.User {
	return models.User{
		Username: userDTO.Username,
		Password: userDTO.Password,
	}
}

func ToUserDTOResponse(user models.User) dto.UserDTOResponse {
	return dto.UserDTOResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
	}
}
