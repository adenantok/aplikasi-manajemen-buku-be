package repositories

import (
	"aplikasi-manajemen-buku-be/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(post *models.Book) (models.Book, error)
	Getbooks() ([]models.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

func (repo *bookRepository) CreateBook(book *models.Book) (models.Book, error) {
	// Menyimpan post ke dalam database
	if err := repo.db.Create(book).Error; err != nil {
		return models.Book{}, err
	}
	return *book, nil
}

func (repo *bookRepository) Getbooks() ([]models.Book, error) {

	var book []models.Book
	if err := repo.db.Find(&book).Error; err != nil {
		return book, err
	}

	return book, nil
}
