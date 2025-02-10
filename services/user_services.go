package services

import (
	"articlehub-be/auth/token"
	"articlehub-be/dto"
	"articlehub-be/mappers"
	"articlehub-be/repositories"
	"articlehub-be/utils"
	"errors"

	"github.com/google/uuid"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) AddUser(userDTO dto.UserDTO) (dto.UserDTOResponse, error) {
	existingUser, err := s.repo.GetUserByEmail(userDTO.Email)
	if err == nil && existingUser.ID != uuid.Nil {
		return dto.UserDTOResponse{}, errors.New("email has been registered")
	}

	user := mappers.ToUser(userDTO)
	passHash, err := utils.HashPassword(user.Password)
	if err != nil {
		return dto.UserDTOResponse{}, err
	}
	user.Password = passHash

	result, err := s.repo.AddUser(user)
	if err != nil {
		return dto.UserDTOResponse{}, err
	}
	return mappers.ToUserDTOResponse(result), nil
}

func (s *UserService) LoginUser(userDTOLogin dto.UserDTOLogin) (dto.UserDTOResponse, string, error) {
	user := mappers.ToUserLogin(userDTOLogin)
	user, err := s.repo.GetUserByEmail(user.Email)
	if err != nil {
		return dto.UserDTOResponse{},"", errors.New("email or password is incorrect")
	}
	if !utils.ComparePassword(user.Password, userDTOLogin.Password) {
		return dto.UserDTOResponse{},"", errors.New("email or password is incorrect")
	}
	token, err := token.GenerateToken(user)
	if err != nil {
		return dto.UserDTOResponse{},"", err
	}

	userDtoRespnse := mappers.ToUserDTOResponse(user)
	return userDtoRespnse, token, nil
}