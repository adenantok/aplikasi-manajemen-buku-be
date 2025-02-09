package routes

import (
	"articlehub-be/auth/middleware"
	"articlehub-be/config"
	"articlehub-be/controllers"
	"articlehub-be/repositories"
	"articlehub-be/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())


	userController := controllers.NewUserController(services.NewUserService(repositories.NewUserRepository(config.DB)))
	router.POST("/register", userController.AddUser)
	router.POST("/login", userController.LoginUser)

	return router
}