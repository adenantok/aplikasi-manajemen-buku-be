package models

import "gorm.io/gorm"

type Book struct {
	ID          int    `gorm:"primaryKey;autoIncrement"`
	UserID      int    `gorm:"not null"`
	Title       string `gorm:"not null"`
	Author      string `gorm:"not null"`
	Description string `gorm:"type:text"`

	User User `gorm:"foreignKey:UserID"` // Relasi dengan tabel User
}

func MigrateBook(db *gorm.DB) error {
	return db.AutoMigrate(&Book{})
}
