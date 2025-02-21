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

	fmt.Println("Connected to database")

	//
	//databaseName := os.Getenv("DATABASE_NAME")
	//if databaseName == "" {
	//	log.Fatal("Database name environment variable not set")
	//}
}
