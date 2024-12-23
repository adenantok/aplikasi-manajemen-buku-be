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

func (s *BookService) GetBooks() ([]dto.BookDTOResponse, error) {

	books, err := s.repo.Getbooks()
	if err != nil {
		return nil, err
	}

	bookDTOResponse := mappers.MapToBooksDTO(books)

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
