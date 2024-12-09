package main

import (
	"aplikasi-manajemen-buku-be/config"
	"aplikasi-manajemen-buku-be/models"
	"aplikasi-manajemen-buku-be/routes"
	"log"
)

func main() {
	config.ConnectDB()
	//config.TestDBConnection()
	models.MigrateUser(config.DB)
	models.MigrateBook(config.DB)

	// Menyiapkan router

	router := routes.SetupRouter()

	// Menjalankan server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server gagal dijalankan: ", err)
	}
}
