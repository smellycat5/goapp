package user

import (
	"Go/initializers"
	"Go/models"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Create(c *gin.Context) {

	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	var existingUser models.User
	data := initializers.DB.Where("email =?", &user.Email).Find(&existingUser)

	if data.RowsAffected > 0 {
		c.JSON(http.StatusConflict, gin.H{"message": "Email already taken."})
		return
	}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			c.JSON(http.StatusConflict, gin.H{"message": "User already exists"})
		} else if errors.Is(result.Error, gorm.ErrForeignKeyViolated) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid foreign key reference"})
		} else if errors.Is(result.Error, gorm.ErrInvalidData) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Invalid data"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Database error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Created!!"})
}

func List(c *gin.Context) {
	var user []models.User

	result := initializers.DB.Find(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No users found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user, "message": "Users list fetched!"})

}

func View(c *gin.Context) {
	var user models.User
	result := initializers.DB.First(&user, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User Not Found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user, "message": "User Found!"})
}

func Delete(c *gin.Context) {
	var user models.User

	result := initializers.DB.Delete(&user, c.Param("id"))

	log.Printf("Rows Effected: %d", result.RowsAffected)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User Not Found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Deleted!"})

}
