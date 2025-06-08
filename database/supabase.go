package database

import (
	"log"
	"os"

	models "github.com/rishichirchi/go-server/model"
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

	err = db.AutoMigrate(
		models.MCP{},
	)

	if err != nil {
		log.Fatalf("Failed to auto migrate models: %v", err)
	}
	log.Println("Auto migration completed successfully")



	DB = db

	return db, nil
}
