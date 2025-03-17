package db

import (
	"fmt"
	"goth-todo/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=goth password=Hockey7232! dbname=goth-todo port=5432"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to the database:", err)
	}
	fmt.Println("✅ Connected to PostgreSQL")
}

// InitializeDB sets up the database connection
func InitializeDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	log.Println("Database connection established")

	// Auto-migrate models
	MigrateDB()
}

// MigrateDB applies automatic migrations
func MigrateDB() {
	err := DB.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatal("Migration failed: ", err)
	}
	log.Println("Database migration complete")
}

// GetDB returns the global database instance
func GetDB() *gorm.DB {
	return DB
}

func CloseDB() {
	sqlDB, err := DB.DB() // Get underlying sql.DB
	if err != nil {
		log.Println("Error getting database connection:", err)
		return
	}
	sqlDB.Close()
	log.Println("Database connection closed")
}
