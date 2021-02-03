package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	err_ "github.com/syx309/training_go/internal/err"
	"net/http"
	"strings"
)

func BasicAuth(handle httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		authHeader := request.Header.Get("authorization")
		if authHeader != "" {
			bearerToken := strings.Split(authHeader, " ")
			if len(bearerToken) == 2 {
				var claims Claims
				token, err := jwt.ParseWithClaims(bearerToken[1], &claims, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, errors.New("JWT Error")
					}
					return []byte(err_.JwtKey), nil
				})
				if err != nil {
					err_.ErrorInternal(writer)
					panic(err)
				}
				if claims, ok := token.Claims.(*Claims); ok && token.Valid {
					request.Header.Set("email", claims.Email)
					handle(writer, request, params)
				} else {
					err_.ErrorUnauthorized(writer)
					panic(errors.New("JWT Error"))
				}
			} else {
				err_.ErrorForbidden(writer)
				panic(errors.New("JWT Error"))
			}
		} else {
			err_.ErrorForbidden(writer)
			panic(errors.New("JWT Error"))
		}
	}
}
