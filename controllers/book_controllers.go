package controllers

import (
	"aplikasi-manajemen-buku-be/dto"
	"aplikasi-manajemen-buku-be/services"
	"aplikasi-manajemen-buku-be/utils"
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
	books, err := controller.service.GetBooks()
	if err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}
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
