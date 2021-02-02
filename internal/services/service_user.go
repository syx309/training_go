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

func RouteUser(writer http.ResponseWriter, request *http.Request, _ httprouter.Params){
	userData := decodeUserData(request)

	query := fmt.Sprintf(`SELECT id, name, email, password FROM users WHERE LOWER(email) = LOWER('%s')`, userData.Email)
	row := helpers.DB.QueryRow(query)

	var user dtos.User
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
			helpers.ErrorNotFound(writer)
			//panic(err)
		} else {
			helpers.ErrorInternal(writer)
			panic(err)
		}
	}

	data, marshallErr := json.Marshal(user)

	if marshallErr != nil {
		helpers.ErrorInternal(writer)
		panic(marshallErr)
	}

	_, _ = writer.Write(data)
}

func decodeUserData(request *http.Request) dtos.UserData {
	var userData dtos.UserData
	userData.Email = request.Header.Get("email")
	return userData
}
