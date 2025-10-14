package services

import (
	"Go/dto"
	"Go/models"
	"errors"
	"gorm.io/gorm"
)

type PostService struct {
	DB *gorm.DB
}

func NewPostService(db *gorm.DB) *PostService {
	return &PostService{DB: db}
}

func (pc *PostService) Create(createPostDto dto.CreatePostDto, userId int) error {

	var post models.Post

	post = models.Post{
		Title:       createPostDto.Title,
		Description: createPostDto.Description,
		UserID:      uint(userId),
	}

	if err := pc.DB.Create(&post).Error; err != nil {
		return err
	}

	return nil
}

func (pc *PostService) View(postId int, userId int) (models.Post, error) {

	var post models.Post
	result := pc.DB.First(&post, postId)

	if result.Error != nil {
		return post, result.Error
	}

	if result.RowsAffected > 0 {
		if post.UserID != uint(userId) {
			return post, errors.New("post not found")
		}
		return post, nil
	}

	if result.RowsAffected == 0 {
		return post, errors.New("post not found")
	}

	return post, nil
}
