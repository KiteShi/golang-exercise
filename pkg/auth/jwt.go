package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var errUninitializedJWT = errors.New("empty JWT key")

var jwtKey []byte

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func InitJWT(key string) error {
	if len(key) == 0 {
		return errUninitializedJWT
	}

	jwtKey = []byte(key)
	return nil
}

func GenerateJWT(username string) (string, error) {
	if err := validateJWTKey(); err != nil {
		return "", err
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenString string) (*Claims, error) {
	if err := validateJWTKey(); err != nil {
		return nil, err
	}

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}

func validateJWTKey() error {
	if len(jwtKey) == 0 {
		return errUninitializedJWT
	}
	return nil
}
