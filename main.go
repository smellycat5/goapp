package main

import (
	"Go/controllers/auth"
	"Go/controllers/post"
	"Go/controllers/user"
	"Go/initializers"
	"Go/middlewares"
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
	postService := services.NewPostService(initializers.DB)

	authController := auth.NewAuthController(userService)
	postController := post.NewPostController(postService)

	//Auth
	r.POST("/auth/register", authController.Register)
	r.POST("/auth/login", authController.Login)

	//Users

	protected := r.Group("/api")
	protected.Use(middlewares.JWTAuthMiddleware())
	{
		protected.POST("/user/create", user.Create)
		protected.POST("/post/create", postController.Create)
	}
	//protected.Use()
	r.GET("/user/:id/view", user.View)
	r.POST("/user/:id/delete", user.Delete)
	r.GET("/user/list", user.List)

	//Posts

	r.Run(":6969")
}

func authMiddleware(c *gin.Context) {

	//TODO: implement authentication

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}
}
