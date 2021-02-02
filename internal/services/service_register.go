package services

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"training_go/cmd/helpers"
)

func RouteRegister(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	registerData := decodeRegister(writer, request)

	query := fmt.Sprintf("INSERT INTO users (name, email, password) VALUES (\"%s\", \"%s\", \"%s\")", registerData.Name, registerData.Email, registerData.Password)
	_, err := helpers.DB.Exec(query)
	if err != nil {
		helpers.ErrorInternal(writer)
	}

	writer.Write([]byte("Register SUCCESS"))
}

func decodeRegister(writer http.ResponseWriter, request *http.Request) Register {
	decoder := json.NewDecoder(request.Body)
	var register Register
	if err := decoder.Decode(&register); err != nil {
		helpers.ErrorInternal(writer)
		panic(err)
	}
	return register
}

type Register struct {
	Email 		string `json:"email"`
	Password 	string `json:"password"`
	Name     	string `json:"name"`
}
