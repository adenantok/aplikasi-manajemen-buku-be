package main

import (
	"articlehub-be/config"
	"articlehub-be/models"
	"articlehub-be/routes"
	"fmt"
	"log"
	// Sesuaikan dengan nama module Anda
)

func main() {
	config.ConnectDB()
	fmt.Println("🎯 Database connection test completed.")

	models.MigrateUser(config.DB)

	router:= routes.SetupRouter()

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server gagal dijalankan: ", err)
	}
}
