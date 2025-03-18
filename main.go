package main

import (
	"Go/controllers/auth"
	"Go/controllers/user"
	"Go/initializers"
	"Go/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	initializers.LoadEnvironmentVariables()
	initializers.ConnectToDatabase()
}

func main() {

	r := gin.Default()

	userService := services.NewUserService(initializers.DB)
	authController := auth.NewAuthController(userService)

	//Auth
	r.POST("/auth/register", authController.Register)
	r.POST("/auth/login", authController.Login)

	//Users

	//protected := r.Group("/")

	//protected.Use()
	r.POST("/user/create", user.Create)
	r.GET("/user/:id/view", user.View)
	r.POST("/user/:id/delete", user.Delete)
	r.GET("/user/list", user.List)

	//Posts

	r.Run(":6969")
}

func authMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}
}
