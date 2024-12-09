package mappers

import (
	"aplikasi-manajemen-buku-be/dto"
	"aplikasi-manajemen-buku-be/models"
)

// MapToUser mengonversi UserDTO ke User model
func MapToUser(userDTO dto.UserDTO) models.User {
	return models.User{
		Username: userDTO.Username,
		Password: userDTO.Password, // Password akan dihash sebelum disimpan di database
	}
}

func MaptoUserDTOResponse(user models.User) dto.UserDTOResponse {
	return dto.UserDTOResponse{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	}
}
