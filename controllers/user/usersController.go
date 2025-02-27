package user

import (
	"Go/initializers"
	"Go/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Create(c *gin.Context) {

	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Created!!"})
}

func View(c *gin.Context) {
	var user models.User
	result := initializers.DB.First(&user, c.Param("id"))

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user, "message": "User Found!"})
}

func Delete(c *gin.Context) {
	var user models.User

	result := initializers.DB.Delete(&user, c.Param("id"))

	log.Printf("Rows Effected: %d", result.RowsAffected)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Deleted!"})

}
