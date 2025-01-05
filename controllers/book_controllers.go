package controllers

import (
	"aplikasi-manajemen-buku-be/dto"
	"aplikasi-manajemen-buku-be/services"
	"aplikasi-manajemen-buku-be/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	service *services.BookService
}

func NewBookController(service *services.BookService) *BookController {
	return &BookController{
		service: service,
	}
}

func (controller *BookController) CreateBook(c *gin.Context) {
	var CreateBookDTO dto.CreateBookDTO

	// Bind the incoming JSON data to postDTO
	if err := c.ShouldBindJSON(&CreateBookDTO); err != nil {
		utils.BadRequestResponse(c, err.Error())
		return
	}

	// // Ambil userID dari konteks
	userID, exists := c.Get("userID")
	if !exists {
		utils.InternalServerErrorResponse(c, "UserID not found in context")
		return
	}

	CreateBookDTO.UserID = userID.(int)

	// Pass the DTO to the PostService to process the logic
	createdPost, err := controller.service.CreateBook(CreateBookDTO)
	if err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}

	// Return success response with created data
	utils.CreatedResponse(c, "Book created successfully", createdPost)
}

func (controller *BookController) GetBooks(c *gin.Context) {
	// Ambil parameter `page` dan `limit` dari query
	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	// Parsing nilai `page` dan `limit`, default jika kosong
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1 // Default halaman pertama
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 0 {
		limit = 0 // Default: ambil semua buku
	}

	// Panggil service untuk mendapatkan buku
	books, err := controller.service.GetBooks(page, limit)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal mendapatkan data buku")
		return
	}

	// Kirim respons JSON
	utils.SuccessResponse(c, "books retrieved successfully", books)
}

func (controller *BookController) GetBookByID(c *gin.Context) {
	bookID := c.Param("id")

	id, err := strconv.Atoi(bookID)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid bookID")
		return
	}

	book, err := controller.service.GetBookByID(id)
	if err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}
	utils.SuccessResponse(c, "book retrieved successfully", book)
}

func (controller *BookController) UpdateBook(c *gin.Context) {
	var bookDTO dto.UpdateBookDTO

	// Bind the incoming JSON data to postDTO
	if err := c.ShouldBindJSON(&bookDTO); err != nil {
		utils.BadRequestResponse(c, err.Error())
		return
	}

	// Pass the DTO to the PostService to process the logic
	updatedBook, err := controller.service.UpdateBook(bookDTO)
	if err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}

	// Return success response with updated data
	utils.SuccessResponse(c, "Book updated successfully", updatedBook)
}

func (controller *BookController) DeleteBook(c *gin.Context) {
	bookID := c.Param("id")

	id, err := strconv.Atoi(bookID)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid bookID")
		return
	}

	err = controller.service.DeleteBook(id)
	if err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}

	utils.SuccessResponse(c, "Book deleted successfully", nil)
}
