package models

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title       string `json:"title"  gorm:"size:100"`
	Description string `json:"description"`
	Image       string `json:"image"` // Stores filename or URL of uploaded image
	UserID      string `json:"user_id" gorm:"size:36;index"`
	User        User   `json:"user"
	gorm:"foreignKey:UserID;references:UserID"`
}

/* gorm.Model auto embedded

type Model struct {
    ID        uint           `gorm:"primaryKey"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}

*/
