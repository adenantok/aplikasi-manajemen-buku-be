package mappers

import (
	"articlehub-be/dto"
	"articlehub-be/models"
)

func ToUser(userDTO dto.UserDTO) models.User {
	return models.User{
		Name:     userDTO.Name,
		Email:    userDTO.Email,
		Password: userDTO.Password,
	}
}

func ToUserLogin(userDTO dto.UserDTOLogin) models.User {
	return models.User{
		Email:    userDTO.Email,
		Password: userDTO.Password,
	}
}

func ToUserDTOResponse(user models.User) dto.UserDTOResponse {
	return dto.UserDTOResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Role:     user.Role,
	}
}

