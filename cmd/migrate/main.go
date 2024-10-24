package main

import (
	"log"
	"os"

	"github.com/Harsh-710/hospital-management/db"
	"github.com/Harsh-710/hospital-management/models"
	"gorm.io/gorm"
)

// applies the database migrations
func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Patient{},
		&models.Appointment{},
	)
	if err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}
	log.Println("Database migrated successfully")
}

func main() {
	db.ConnectDB()

	database := db.GetDB()	// gets the gorm database instance

	RunMigrations(database)

	log.Println("Migration complete, exiting...")
	os.Exit(0)
}
