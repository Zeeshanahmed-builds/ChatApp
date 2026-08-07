package db

import (
	"fmt"
	"log"
	"os"
	"github.com/Zeeshanahmed-builds/ChatApp/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {

	// Load .env (optional - will use environment variables if not found)
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using environment variables from ConfigMap/Secret")
	}

	// Read environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Build DSN
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
		host,
		port,
		user,
		password,
		dbName,
	)

	// Connect
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	// Auto migrate
	err = db.AutoMigrate(
		&models.User{},
		&models.Message{},
	)
	if err != nil {
		return nil, fmt.Errorf("auto migration failed: %v", err)
	}

	log.Println("Database connected successfully!")

	return db, nil
}
