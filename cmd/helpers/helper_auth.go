package helpers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
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
					return []byte(jwtKey), nil
				})
				if err != nil {
					ErrorInternal(writer)
					panic(err)
				}
				if claims, ok := token.Claims.(*Claims); ok && token.Valid {
					request.Header.Set("email", claims.Email)
					handle(writer, request, params)
				} else {
					ErrorUnauthorized(writer)
					panic(errors.New("JWT Error"))
				}
			} else {
				ErrorForbidden(writer)
				panic(errors.New("JWT Error"))
			}
		} else {
			ErrorForbidden(writer)
			panic(errors.New("JWT Error"))
		}
	}
}
