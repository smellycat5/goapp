package post

import (
	"Go/dto"
	"Go/models"
	"Go/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostController struct {
	PostService *services.PostService
}

func NewPostController(postService *services.PostService) *PostController {
	return &PostController{PostService: postService}
}

func (pc *PostController) Create(c *gin.Context) {
	var post models.Post

	if err := c.ShouldBind(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	postDto := dto.CreatePostDto{
		Title:       post.Title,
		Description: post.Description,
	}

	err := pc.PostService.Create(postDto, 5)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post Created!!"})

}

func (pc *PostController) View(c *gin.Context) {
	var post models.Post
	pc.PostService.View(c)
}

func (pc *PostController) Delete(c *gin.Context) {}
