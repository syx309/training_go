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

func RouteGetItemByName(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	responseBody := decodeItemData(writer, request)

	query := `SELECT items.id, user_id, app_name, app_email, app_password 
								FROM items JOIN users 
								ON items.user_id = users.id
								WHERE users.email = $1 
								AND LOWER(items.app_name) = LOWER($2)`
	row := datastore.DB.QueryRow(query, responseBody.Email, responseBody.AppName)

	var item dtos.Item
	err := row.Scan(&item.Id, &item.User_id, &item.App_name, &item.App_email, &item.App_password)

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
	data, marshallErr := json.Marshal(item)

	if marshallErr != nil {
		err2.ErrorInternal(writer)
		panic(marshallErr)
	}

	_, _ = writer.Write(data)
}

func decodeItemData(writer http.ResponseWriter, request *http.Request) GetItemData {
	decoder := json.NewDecoder(request.Body)
	var itemData GetItemData
	if err := decoder.Decode(&itemData); err != nil {
		err2.ErrorInternal(writer)
		//panic(err)
	}
	return itemData
}

type GetItemData struct {
	Email   string `json:"email"`
	AppName string `json:"appName"`
}
