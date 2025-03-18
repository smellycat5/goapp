package main

import (
	"Go/initializers"
	"Go/models"
	"log"
)

func init() {
	initializers.LoadEnvironmentVariables()
	initializers.ConnectToDatabase()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.User{}, &models.Post{})
	if err != nil {
		log.Fatal("Cannot Migrate database")
	} else {
		log.Println("Database migrated")
	}

}
