package main

import (
	"Go/initializers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	initializers.LoadEnvironmentVariables()
}

func main() {

	x := 1
	r := gin.Default()

	message := ""
	if x >= 1 {
		message = "hello, chunan"
	} else {
		message = "hello dude"
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": message})
	})

	r.Run(":6969")
}
