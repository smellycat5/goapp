package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title       string    `json:"title" binding:"required"`
	UUID        uuid.UUID `gorm:"type:char(36);uniqueIndex" json:"uuid" binding:"required"`
	Description string    `json:"description" binding:"required"`
	UserId      int       `json:"userId" binding:"required"`
}

func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	p.UUID = uuid.New()
	return
}
