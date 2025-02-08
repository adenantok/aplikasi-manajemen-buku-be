package services

import (
	"articlehub-be/auth/token"
	"articlehub-be/dto"
	"articlehub-be/mappers"
	"articlehub-be/repositories"
	"articlehub-be/utils"
	"errors"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) AddUser(userDTo dto.UserDTO) (dto.UserDTOResponse, error) {
	user := mappers.ToUser(userDTo)
	user.Password, _ = utils.HashPassword(user.Password)
	user, err := s.repo.AddUser(user)
	return mappers.ToUserDTOResponse(user), err
}

func (s *UserService) LoginUser(userDTOLogin dto.UserDTOLogin) (dto.UserDTOResponse, string, error) {
	user := mappers.ToUserLogin(userDTOLogin)
	user, err := s.repo.GetUserByUsername(user.Username)
	if err != nil {
		return dto.UserDTOResponse{},"", err
	}
	if !utils.ComparePassword(user.Password, userDTOLogin.Password) {
		return dto.UserDTOResponse{},"", errors.New("invalid username or password")
	}
	token, err := token.GenerateToken(user)
	if err != nil {
		return dto.UserDTOResponse{},"", err
	}

	userDtoRespnse := mappers.ToUserDTOResponse(user)
	return userDtoRespnse, token, nil
}