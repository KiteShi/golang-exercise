package db

import (
	"fmt"
	"log"
	"time"

	"github.com/KiteShi/golang-exercise/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	maxConnectionAttempts = 3
	waitTimeSec           = 5
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func InitDB(cfg Config) {
	var attempt int
	for attempt < maxConnectionAttempts {
		attempt++
		log.Printf("Attempting to connect to the database (attempt %d/%d)...", attempt, maxConnectionAttempts)
		db, err := gorm.Open(postgres.Open(getDSN(cfg)), &gorm.Config{})
		if err != nil {
			log.Printf("failed to connect to the database: %v", err)
			time.Sleep(waitTimeSec * time.Second) // Wait before retrying
			continue
		}

		log.Print("database connection established")
		DB = db
		// Auto migrate database schema
		migrateDB()
		break
	}

	if attempt >= maxConnectionAttempts {
		log.Fatalf("Failed to connect to the database after %d attempts", maxConnectionAttempts)
	}
}

func CloseDB() {
	db, err := DB.DB()
	if err != nil {
		log.Fatalf("failed to get DB instance: %v", err)
	}
	db.Close()
}

func migrateDB() {
	if err := DB.AutoMigrate(&models.Company{}); err != nil {
		log.Fatalf("failed to migrate database schema: %v", err)
	}
}

func getDSN(db Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		db.Host, db.Port, db.User, db.Password, db.DBName)
}
