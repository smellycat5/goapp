package post

import (
	"Go/dto"
	"Go/models"
	"Go/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func (pc *PostController) View(c *gin.Context) {
	uuid := c.Param("uuid")

	if uuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "uuid cannot be empty"})
		return
	}

	post, err := pc.PostService.View(uuid)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Post fetched Successfully!", "data": post})
	}
}

func (pc *PostController) Delete(c *gin.Context) {
	uuid := c.Param("uuid")

	if uuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "uuid cannot be empty"})
		return
	}

	err := pc.PostService.Delete(uuid)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Post Deleted Successfully!"})
	}

}
