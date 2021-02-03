package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/syx309/training_go/cmd/datastore"
	"github.com/syx309/training_go/internal/dtos"
	err2 "github.com/syx309/training_go/internal/err"
	"net/http"
)

func RouteUser(writer http.ResponseWriter, request *http.Request, _ httprouter.Params){
	userData := decodeUserData(request)

	query := fmt.Sprintf(`SELECT id, name, email, password FROM users WHERE LOWER(email) = LOWER('%s')`, userData.Email)
	row := datastore.DB.QueryRow(query)

	var user dtos.User
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
			err2.ErrorNotFound(writer)
			//panic(err)
		} else {
			err2.ErrorInternal(writer)
			panic(err)
		}
	}

	data, marshallErr := json.Marshal(user)

	if marshallErr != nil {
		err2.ErrorInternal(writer)
		panic(marshallErr)
	}

	_, _ = writer.Write(data)
}

func decodeUserData(request *http.Request) dtos.UserData {
	var userData dtos.UserData
	userData.Email = request.Header.Get("email")
	return userData
}
