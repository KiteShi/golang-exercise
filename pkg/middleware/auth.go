package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/KiteShi/golang-exercise/pkg/auth"
)

type contextKey string

const UserContextKey = contextKey("user")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := auth.ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			log.Printf("Invalid token: %v", err)
			return
		}
		log.Printf("User %s authenticated", claims.Username)

		ctx := context.WithValue(r.Context(), UserContextKey, claims.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
