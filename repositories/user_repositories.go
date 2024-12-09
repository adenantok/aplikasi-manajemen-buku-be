package repositories

import (
	"aplikasi-manajemen-buku-be/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByUsername(username string) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	if err := repo.db.Where("username = ?", username).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
