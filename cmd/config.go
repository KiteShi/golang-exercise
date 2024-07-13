package cmd

import (
	"os"

	"github.com/KiteShi/golang-exercise/pkg/auth"
	"github.com/KiteShi/golang-exercise/pkg/db"
)

func InitDBConfig() db.Config {
	var cfg db.Config
	cfg.Host = os.Getenv("DB_HOST")
	cfg.Port = os.Getenv("DB_PORT")
	cfg.User = os.Getenv("DB_USER")
	cfg.Password = os.Getenv("DB_PASSWORD")
	cfg.DBName = os.Getenv("DB_NAME")

	return cfg
}

func InitJWT() error {
	jwtKey := os.Getenv("JWT_KEY")
	return auth.InitJWT(jwtKey)
}

func InitAdmins() error {
	adminsEnv := os.Getenv("ADMINS")
	return auth.InitAdmins(adminsEnv)
}
