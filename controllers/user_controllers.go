package controllers

import (
	"articlehub-be/dto"
	"articlehub-be/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (controller *UserController) AddUser(c *gin.Context) {
	var userDTO dto.UserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		
	}

	user, err := controller.service.AddUser(userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (controller *UserController) LoginUser(c *gin.Context) {
	var userDTO dto.UserDTOLogin
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user,token,  err := controller.service.LoginUser(userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	responseData := gin.H{
		"user": user,
		"token": token,
	}

	c.JSON(http.StatusOK, responseData)

}