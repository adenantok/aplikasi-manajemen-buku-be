package repositories

import (
	"aplikasi-manajemen-buku-be/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(post *models.Book) (models.Book, error)
	GetAllBooks() ([]models.Book, error)
	GetBookByID(id int) (models.Book, error)
	UpdateBook(book *models.Book) (models.Book, error)
	DeleteBook(id int) error
	GetBooksPaginated(limit, offset int) ([]models.Book, error)
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

func (repo *bookRepository) GetAllBooks() ([]models.Book, error) {

	var book []models.Book
	if err := repo.db.Order("id DESC").Find(&book).Error; err != nil {
		return book, err
	}

	return book, nil
}

func (repo *bookRepository) GetBookByID(id int) (models.Book, error) {

	var book models.Book

	if err := repo.db.First(&book, id).Error; err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (repo *bookRepository) UpdateBook(book *models.Book) (models.Book, error) {

	if err := repo.db.Save(&book).Error; err != nil {
		return models.Book{}, err
	}

	return *book, nil
}

func (repo *bookRepository) DeleteBook(id int) error {
	return repo.db.Delete(&models.Book{}, id).Error
}

func (repo *bookRepository) GetBooksPaginated(limit, offset int) ([]models.Book, error) {
	var books []models.Book
	result := repo.db.Order("id DESC").Limit(limit).Offset(offset).Find(&books)
	return books, result.Error
}
