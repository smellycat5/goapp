package user

import (
	"Go/initializers"
	"Go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(c *gin.Context) {
	//get Data from request and create

	user := models.User{Name: c.PostForm("name"), Email: c.PostForm("email")}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Created!"})
}

func View(c *gin.Context) {
	var user models.User
	initializers.DB.First(&user, c.Param("id"))

	c.JSON(http.StatusOK, gin.H{"data": user, "message": "User Found!"})
}

func Delete(c *gin.Context) {
	initializers.DB.Delete(models.User{}, c.Param("id"))

	c.JSON(http.StatusOK, gin.H{"message": "User Deleted!"})

}
