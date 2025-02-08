package routes

import (
	"articlehub-be/config"
	"articlehub-be/controllers"
	"articlehub-be/repositories"
	"articlehub-be/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	userController := controllers.NewUserController(services.NewUserService(repositories.NewUserRepository(config.DB)))
	router.POST("/users", userController.AddUser)
	router.POST("/login", userController.LoginUser)

	return router
}