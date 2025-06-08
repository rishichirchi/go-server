package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitSupabase() (*gorm.DB, error) {
	dsn := os.Getenv("SUPABASE_URL")

	if dsn == "" {
		log.Fatal("SUPABASE_URL is not set in .env file")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to Supabase: %v", err)
	}
	log.Println("Connected to Supabase successfully")

	DB = db

	return db, nil
}
