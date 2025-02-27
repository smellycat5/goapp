package main

import (
	"Go/controllers/user"
	"Go/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvironmentVariables()
	initializers.ConnectToDatabase()
}

func main() {

	r := gin.Default()

	r.POST("/user/create", user.Create)
	r.GET("/user/:id/view", user.View)
	r.POST("/user/:id/delete/", user.Delete)

	r.Run(":6969")
}
