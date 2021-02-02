package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"training_go/cmd/helpers"
	"training_go/internal/dtos"
)

func RouteLogin(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	loginData := decodeLogin(writer, request)

	query := fmt.Sprintf(`SELECT id, name, email, password FROM users WHERE LOWER(email) = LOWER('%s')`, loginData.Email)
	row := helpers.DB.QueryRow(query)

	var user dtos.User
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
			helpers.ErrorNotFound(writer)
			panic(err)
		} else {
			helpers.ErrorInternal(writer)
			panic(err)
		}
	}

	tokenString := helpers.GenerateJWT(user.Email)
	writer.Write([]byte("Halo " + user.Email + ", token kamu: " + tokenString))
}

func decodeLogin(writer http.ResponseWriter, request *http.Request) Login {
	decoder := json.NewDecoder(request.Body)
	var login Login
	if err := decoder.Decode(&login); err != nil {
		helpers.ErrorInternal(writer)
		panic(err)
	}
	return login
}

type Login struct {
	Email 		string `json:"email"`
	Password 	string `json:"password"`
}
