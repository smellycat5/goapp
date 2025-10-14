package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	UserId      int    `json:"userId" binding:"required"`
}
