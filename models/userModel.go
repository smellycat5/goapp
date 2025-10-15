package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" gorm:"uniqueIndex" binding:"required,email"`
	Password string `json:"-"`
	IsAdmin  int16  `gorm:"type:smallint;default:0"`
	Posts    []Post `json:"posts"`
}
