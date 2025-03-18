package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title       string `json:"name" binding:"required"`
	Description string `json:"email" gorm:"unique" binding:"required,email"`
	UserID      uint
}
