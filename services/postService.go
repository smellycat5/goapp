package services

import (
	"Go/dto"
	"Go/models"
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
		UserId:      userId,
	}

	if err := pc.DB.Create(&post).Error; err != nil {
		return err
	}

	return nil
}

func (pc *PostService) View(postUuid string) (models.Post, error) {

	var post models.Post
	result := pc.DB.Where("uuid = ?", postUuid).First(&post)

	if result.Error != nil {
		return models.Post{}, result.Error
	}

	return post, nil
}

func (pc *PostService) Delete(postUuid string) error {
	var post models.Post
	result := pc.DB.Where("uuid = ?", postUuid).First(&post)

	if result.Error != nil {
		return result.Error
	}

	if err := pc.DB.Delete(&post).Error; err != nil {
		return err
	}

	return nil
}

func (pc *PostService) Update(updatePostDto dto.UpdatePostDto) error {

	var post models.Post

	result := pc.DB.Where("uuid=?", updatePostDto.UUID).First(&post)

	if result.Error != nil {
		return result.Error
	}

	post.Title = updatePostDto.Title
	post.Description = updatePostDto.Description

	if err := pc.DB.Save(&post).Error; err != nil {
		return err
	}
	return nil
}
