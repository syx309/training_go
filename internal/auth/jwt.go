package auth

import (
	"github.com/dgrijalva/jwt-go"
	err_ "github.com/syx309/training_go/internal/err"
	"time"
)

func GenerateJWT(email string) string {
	expirationTime := time.Now().Add(48 * time.Hour)
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString([]byte(err_.JwtKey))
	if err != nil {
		panic(err)
	}
	return tokenString
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

