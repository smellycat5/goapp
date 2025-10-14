package dto

type CreatePostDto struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}
