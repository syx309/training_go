package helpers

import (
	"github.com/dgrijalva/jwt-go"
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
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		panic(err)
	}
	return tokenString
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

