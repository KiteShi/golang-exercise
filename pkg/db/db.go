package db

import (
	"fmt"
	"log"

	"github.com/KiteShi/golang-exercise/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func InitDB(cfg DBConfig) {
	db, err := gorm.Open(postgres.Open(getDSN(cfg)), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	log.Print("database connection established")

	DB = db

	// Auto migrate database schema
	migrateDB()
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

func getDSN(db DBConfig) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		db.Host, db.Port, db.User, db.Password, db.DBName)
}
