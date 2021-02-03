package services

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/syx309/training_go/cmd/datastore"
	err2 "github.com/syx309/training_go/internal/err"
	"net/http"
)

func RouteUpdateItem(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	responseBody := decodeUpdateData(writer, request)

	query := `UPDATE items SET app_email = $1, app_password = $2 WHERE user_id = $3 AND LOWER(app_name) = LOWER($4)`

	_, err := datastore.DB.Exec(query, responseBody.AppEmail, responseBody.AppPassword, responseBody.UserId, responseBody.AppName)
	if err != nil {
		err2.ErrorInternal(writer)
		panic(err)
	}
	writer.Write([]byte("Change Email and Password SUCCESS"))
}

func decodeUpdateData(writer http.ResponseWriter, request *http.Request) UpdateItemData {
	decoder := json.NewDecoder(request.Body)
	var itemData UpdateItemData
	if err := decoder.Decode(&itemData); err != nil {
		err2.ErrorInternal(writer)
		panic(err)
	}
	return itemData
}

type UpdateItemData struct {
	UserId   		string `json:"userID"`
	AppName  		string `json:"appName"`
	AppEmail  		string `json:"appEmail"`
	AppPassword  	string `json:"appPassword"`
}
