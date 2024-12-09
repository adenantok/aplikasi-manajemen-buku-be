package dto

// BookDTO adalah struktur untuk mentransfer data Book
// type BookDTO struct {
// 	ID          int    `json:"id"`
// 	Title       string `json:"title"`
// 	Author      string `json:"author"`
// 	Description string `json:"description"`
// 	UserID      int    `json:"user_id"`
// }

type CreateBookDTO struct {
	UserID      int    `json:"user_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type BookDTOResponse struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

type BookDTOs struct {
	Books []BookDTOResponse `json:"books"`
}
