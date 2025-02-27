package initializers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	databaseUrl := os.Getenv("DSN")
	if databaseUrl == "" {
		log.Fatal("Database URL environment variable not set")
	}

	DB, err = gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	fmt.Println("*************************")
	fmt.Println("Connected to database :)")
	fmt.Println("*************************")

}
