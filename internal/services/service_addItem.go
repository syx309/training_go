package services

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/syx309/training_go/cmd/datastore"
	err2 "github.com/syx309/training_go/internal/err"
	"net/http"
)

func RouteAddItem(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	responseBody := decodeAddData(writer, request)

	query := `INSERT INTO items(user_id, app_name, app_email, app_password) 
				VALUES($1, $2, $3, $4)`

	_, err := datastore.DB.Exec(query, responseBody.UserId, responseBody.AppName, responseBody.AppEmail, responseBody.AppPassword)
	if err != nil {
		err2.ErrorInternal(writer)
		panic(err)
	}
	writer.Write([]byte("Insert item SUCCESS"))
}

func decodeAddData(writer http.ResponseWriter, request *http.Request) AddItemData {
	decoder := json.NewDecoder(request.Body)
	var itemData AddItemData
	if err := decoder.Decode(&itemData); err != nil {
		err2.ErrorInternal(writer)
		panic(err)
	}
	return itemData
}

type AddItemData struct {
	UserId   		string `json:"userID"`
	AppName  		string `json:"appName"`
	AppEmail  		string `json:"appEmail"`
	AppPassword  	string `json:"appPassword"`
}
