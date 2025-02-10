package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)
type User struct {
    ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`  
    Name      string    `gorm:"not null"`
    Email     string    `gorm:"uniqueIndex"`
    Password  string    `gorm:"not null"` // Disimpan dalam bentuk hash
    Role      string    `gorm:"index;default:user"` // admin / user
    CreatedAt time.Time
}

func MigrateUser(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}