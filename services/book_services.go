package services

import (
	"aplikasi-manajemen-buku-be/dto"
	"aplikasi-manajemen-buku-be/mappers"
	"aplikasi-manajemen-buku-be/repositories"
)

// BookService mengatur logika bisnis terkait buku
type BookService struct {
	repo repositories.BookRepository
}

// NewBookService mengembalikan instance BookService dengan dependency yang di-inject
func NewBookService(repo repositories.BookRepository) *BookService {
	return &BookService{
		repo: repo,
	}
}

func (s *BookService) CreateBook(createBookDTO dto.CreateBookDTO) (dto.BookDTOResponse, error) {

	book := mappers.MapToBook(createBookDTO)

	// Menyimpan buku ke database melalui repository
	createdBook, err := s.repo.CreateBook(&book)
	if err != nil {
		return dto.BookDTOResponse{}, err
	}

	bookDTOResponse := mappers.MapToBookDTOResponse(createdBook)

	return bookDTOResponse, nil
}

func (s *BookService) GetBookByID(id int) (dto.BookDTOResponse, error) {

	book, err := s.repo.GetBookByID(id)
	if err != nil {
		return dto.BookDTOResponse{}, err
	}

	bookDTOResponse := mappers.MapToBookDTOResponse(book)

	return bookDTOResponse, nil
}

func (s *BookService) UpdateBook(updateBookDTO dto.UpdateBookDTO) (dto.BookDTOResponse, error) {

	book := mappers.MapToBookUpdateDTO(updateBookDTO)

	existingBook, err := s.repo.GetBookByID(book.ID)
	if err != nil {
		return dto.BookDTOResponse{}, err
	}

	if book.Title != "" {
		existingBook.Title = book.Title
	}
	if book.Author != "" {
		existingBook.Author = book.Author
	}
	if book.Description != "" {
		existingBook.Description = book.Description
	}

	updatedBook, err := s.repo.UpdateBook(&existingBook)
	if err != nil {
		return dto.BookDTOResponse{}, err
	}

	bookDTOResponse := mappers.MapToBookDTOResponse(updatedBook)

	return bookDTOResponse, nil
}

func (s *BookService) DeleteBook(id int) error {
	return s.repo.DeleteBook(id)
}

func (s *BookService) GetBooks(page, limit int) ([]dto.BookDTOResponse, error) {
	if limit == 0 { // Jika limit tidak ditentukan, ambil semua data
		books, err := s.repo.GetAllBooks()
		if err != nil {
			return nil, err
		}
		// Pemetaan ke DTO
		return mappers.MapToBooksDTO(books), nil
	}

	// Hitung offset berdasarkan halaman dan limit
	offset := (page - 1) * limit
	books, err := s.repo.GetBooksPaginated(limit, offset)
	if err != nil {
		return nil, err
	}

	// Pemetaan ke DTO
	return mappers.MapToBooksDTO(books), nil
}
