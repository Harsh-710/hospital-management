package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB initializes the database connection and stores it in the DB variable
func ConnectDB() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbString := os.Getenv("POSTGRES_URI")

	// Import the PostgreSQL connection string
	dsn := dbString

	// Connect to the PostgreSQL database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	log.Println("Database connection established")
	DB = db
}

// GetDB returns the database connection instance
func GetDB() *gorm.DB {
	return DB
}
