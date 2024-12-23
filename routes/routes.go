package routes

import (
	"aplikasi-manajemen-buku-be/auth/midleware"
	"aplikasi-manajemen-buku-be/config"
	"aplikasi-manajemen-buku-be/controllers"
	"aplikasi-manajemen-buku-be/repositories"
	"aplikasi-manajemen-buku-be/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	//config.ConnectDB()

	// Membuat instance dari router Gin
	router := gin.Default()

	// Membuat instance dari repository, service, dan controller
	userRepo := repositories.NewUserRepository(config.DB) // Membuat instance UserRepository
	userService := services.NewUserService(userRepo)      // Membuat instance UserService
	userController := controllers.NewUserController(userService)

	bookRepo := repositories.NewBookRepository(config.DB) // Membuat instance UserRepository
	bookService := services.NewBookService(bookRepo)      // Membuat instance UserService
	bookController := controllers.NewBookController(bookService)

	router.Use(midleware.CORSMiddleware())
	router.POST("/login", userController.LoginUser)
	protected := router.Group("/")
	protected.Use(midleware.AuthMiddleware())
	{
		protected.POST("/books", bookController.CreateBook)
		protected.GET("/books", bookController.GetBooks)
		protected.GET("/books/:id", bookController.GetBookByID)
		protected.PUT("/books", bookController.UpdateBook)
		protected.DELETE("/books/:id", bookController.DeleteBook)
	}

	// router.POST("/login", userController.LoginUser)
	// router.POST("/books", bookController.CreateBook)
	// router.GET("/books", bookController.GetBooks)

	return router
}
