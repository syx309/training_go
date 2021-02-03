package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/syx309/training_go/cmd/datastore"
	"github.com/syx309/training_go/internal/auth"
	"github.com/syx309/training_go/internal/dtos"
	err2 "github.com/syx309/training_go/internal/err"
	"net/http"
)

func RouteLogin(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	loginData := decodeLogin(writer, request)

	query := fmt.Sprintf(`SELECT id, name, email, password FROM users WHERE LOWER(email) = LOWER('%s')`, loginData.Email)
	row := datastore.DB.QueryRow(query)

	var user dtos.User
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
			err2.ErrorNotFound(writer)
			panic(err)
		} else {
			err2.ErrorInternal(writer)
			panic(err)
		}
	}

	tokenString := auth.GenerateJWT(user.Email)
	writer.Write([]byte("Halo " + user.Email + ", token kamu: " + tokenString))
}

func decodeLogin(writer http.ResponseWriter, request *http.Request) Login {
	decoder := json.NewDecoder(request.Body)
	var login Login
	if err := decoder.Decode(&login); err != nil {
		err2.ErrorInternal(writer)
		panic(err)
	}
	return login
}

type Login struct {
	Email 		string `json:"email"`
	Password 	string `json:"password"`
}
