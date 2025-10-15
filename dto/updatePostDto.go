package dto

type UpdatePostDto struct {
	UUID        string `json:"uuid" binding:"required"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
