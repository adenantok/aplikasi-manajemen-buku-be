package dto

// UserDTO adalah struktur untuk menerima data user dari client
type UserDTO struct {
	//ID       int    `json:"id"`
	Username string `json:"username" binding:"required"` // Username diperlukan
	Password string `json:"password" binding:"required"`
	//Role     string `json:"role" `
}

type UserDTOResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"` // Username diperlukan
	Role     string `json:"role" `
}
