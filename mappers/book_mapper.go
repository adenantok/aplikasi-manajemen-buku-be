package mappers

import (
	"aplikasi-manajemen-buku-be/dto"
	"aplikasi-manajemen-buku-be/models"
)

func MapToBook(createBookDTO dto.CreateBookDTO) models.Book {
	return models.Book{
		UserID:      createBookDTO.UserID,
		Title:       createBookDTO.Title,
		Author:      createBookDTO.Author,
		Description: createBookDTO.Description,
	}
}

// MapBookToBookResponseDTO memetakan model Book ke BookResponseDTO
func MapToBookDTOResponse(book models.Book) dto.BookDTOResponse {
	return dto.BookDTOResponse{
		ID:          book.ID,
		UserID:      book.UserID,
		Title:       book.Title,
		Author:      book.Author,
		Description: book.Description,
	}
}

func MapToBooksDTO(books []models.Book) []dto.BookDTOResponse {
	var bookDTOs []dto.BookDTOResponse
	for _, book := range books {
		bookDTOs = append(bookDTOs, MapToBookDTOResponse(book))
	}
	return bookDTOs
}

func MapToBookUpdateDTO(book dto.UpdateBookDTO) models.Book {
	return models.Book{
		ID:          book.ID,
		Title:       book.Title,
		Author:      book.Author,
		Description: book.Description,
	}
}
