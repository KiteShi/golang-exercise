package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KiteShi/golang-exercise/pkg/auth"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LogIn(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Printf("Error decoding login request: %v", err)
		return
	}

	validCreds, err := auth.Authenticate(creds.Username, creds.Password)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("Error authenticate admin: %v", err)
		return

	}
	if !validCreds {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		log.Printf("Failed login attempt for username %s", creds.Username)
		return
	}

	token, err := auth.GenerateJWT(creds.Username)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		log.Printf("Error generating JWT for username %s: %v", creds.Username, err)
		return
	}
	log.Printf("User %s logged in successfully", creds.Username)

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
