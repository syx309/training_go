package user_services

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/syx309/training_go/cmd/datastore"
	"github.com/syx309/training_go/internal/dtos"
	err2 "github.com/syx309/training_go/internal/err"
	"net/http"
)

func RouteRegister(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	registerData := decodeRegister(writer, request)

	query := fmt.Sprintf("INSERT INTO users (name, email, password) VALUES (\"%s\", \"%s\", \"%s\")", registerData.Name, registerData.Email, registerData.Password)
	_, err := datastore.DB.Exec(query)
	if err != nil {
		err2.ErrorInternal(writer)
	}

	writer.Write([]byte("Register SUCCESS"))
}

func decodeRegister(writer http.ResponseWriter, request *http.Request) dtos.Register {
	decoder := json.NewDecoder(request.Body)
	var register dtos.Register
	if err := decoder.Decode(&register); err != nil {
		err2.ErrorInternal(writer)
		panic(err)
	}
	return register
}
