package services

import (
	"aplikasi-manajemen-buku-be/auth/token"
	"aplikasi-manajemen-buku-be/dto"
	"aplikasi-manajemen-buku-be/mappers"
	"aplikasi-manajemen-buku-be/repositories"
	"aplikasi-manajemen-buku-be/utils"
	"errors"
)

type UserService struct {
	repo repositories.UserRepository
}

// NewUserService membuat instance baru dari userService
func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// LoginUser memverifikasi kredensial pengguna
func (s *UserService) LoginUser(userDTO dto.UserDTO) (dto.UserDTOResponse, string, error) {
	// Mengonversi UserDTO ke dalam model User
	user := mappers.MapToUser(userDTO)

	// Cari pengguna berdasarkan username melalui repository
	user, err := s.repo.GetUserByUsername(user.Username)
	if err != nil {
		return dto.UserDTOResponse{}, "", errors.New("username tidak ditemukan") // Kembalikan error jika user tidak ditemukan
	}

	// Verifikasi password menggunakan auth/service
	if !utils.ComparePassword(user.Password, userDTO.Password) {
		return dto.UserDTOResponse{}, "", errors.New("invalid username or password")
	}

	// Generate JWT token jika login berhasil
	token, err := token.GenerateToken(user)
	if err != nil {
		return dto.UserDTOResponse{}, "", err
	}

	userDTOResponse := mappers.MaptoUserDTOResponse(user)
	return userDTOResponse, token, nil // Kembalikan user jika berhasil login
}
