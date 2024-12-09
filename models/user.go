package models

import "gorm.io/gorm"

type User struct {
	ID       int    `gorm:"primaryKey;autoIncrement"` // ID adalah primary key dan auto increment
	Username string `gorm:"unique;not null"`          // Username harus unik dan tidak boleh null
	Password string `gorm:"not null"`                 // Password tidak boleh null
	Role     string `gorm:"not null"`                 // Role harus berupa string (admin atau user)
}

func MigrateUser(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
