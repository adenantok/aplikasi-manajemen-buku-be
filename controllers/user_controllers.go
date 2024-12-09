package controllers

import (
	"aplikasi-manajemen-buku-be/dto"
	"aplikasi-manajemen-buku-be/services"
	"aplikasi-manajemen-buku-be/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

// NewUserController membuat instance baru dari userController
func NewUserController(service *services.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (controller *UserController) LoginUser(c *gin.Context) {
	var loginDTO dto.UserDTO

	// Bind data dari request body ke LoginDTO
	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		// Jika data yang diterima tidak valid (misalnya field yang wajib tidak ada)
		utils.BadRequestResponse(c, err.Error())
		return
	}

	// Panggil service untuk melakukan login
	user, token, err := controller.service.LoginUser(loginDTO)
	if err != nil {
		// Jika terjadi error pada service (misalnya username atau password salah)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	responseData := gin.H{
		"user":  user,
		"token": token,
	}
	utils.SuccessResponse(c, "login successful", responseData)
}
