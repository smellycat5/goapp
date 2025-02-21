package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt int
}
