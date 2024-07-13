package main

import (
	"log"
	"net/http"
	"os"

	"github.com/KiteShi/golang-exercise/cmd"
	"github.com/KiteShi/golang-exercise/pkg/db"
)

func main() {
	db.InitDB(cmd.InitDBConfig())
	defer db.CloseDB()

	if err := cmd.InitAdmins(); err != nil {
		log.Fatalf("failed initialize administrator credentials: %v", err)
	}

	if err := cmd.InitJWT(); err != nil {
		log.Fatalf("failed initialize jwt key: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable not set")
	}

	log.Printf("Starting server on port %s...", port)
	err := http.ListenAndServe(":"+port, cmd.GetRouter())
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
