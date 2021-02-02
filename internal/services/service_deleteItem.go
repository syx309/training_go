package services

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"training_go/cmd/helpers"
)

func RouteDeleteItem(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	responseBody := decodeDeleteData(writer, request)

	query := `DELETE FROM items WHERE user_id = $1 AND items.id = $2`

	_, err := helpers.DB.Exec(query, responseBody.UserId, responseBody.ItemId)
	if err != nil {
		helpers.ErrorInternal(writer)
		panic(err)
	}
	writer.Write([]byte("Delete item SUCCESS"))
}

func decodeDeleteData(writer http.ResponseWriter, request *http.Request) DeleteItemData {
	decoder := json.NewDecoder(request.Body)
	var itemData DeleteItemData
	if err := decoder.Decode(&itemData); err != nil {
		helpers.ErrorInternal(writer)
		panic(err)
	}
	return itemData
}

type DeleteItemData struct {
	UserId   		string `json:"userID"`
	ItemId  		string `json:"itemID"`
}
