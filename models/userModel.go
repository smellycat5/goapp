package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" gorm:"uniqueIndex" binding:"required,email"`
	Password string `json:"-"`
	Posts    []Post `json:"posts"`
}
