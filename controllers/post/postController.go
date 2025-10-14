package post

import (
	"Go/dto"
	"Go/models"
	"Go/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type PostController struct {
	PostService *services.PostService
}

func NewPostController(postService *services.PostService) *PostController {
	return &PostController{PostService: postService}
}

func (pc *PostController) Create(c *gin.Context) {
	var postDto dto.CreatePostDto

	if err := c.ShouldBindJSON(&postDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userInterface, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}
	log.Printf("*************************************")
	log.Printf("%v\n\n", userInterface)
	user, ok := userInterface.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User type invalid"})
		return
	}
	log.Printf("****************   %v\n\n", user)

	err := pc.PostService.Create(postDto, int(user.ID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Post Created!!"})

	}

}

//
//func (pc *PostController) View(c *gin.Context) {
//	var post models.Post
//	pc.PostService.View(c)
//}

func (pc *PostController) Delete(c *gin.Context) {}
