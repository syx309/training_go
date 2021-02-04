package item_services

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/syx309/training_go/cmd/datastore"
	"github.com/syx309/training_go/internal/dtos"
	err2 "github.com/syx309/training_go/internal/err"
	"net/http"
)

func RouteDeleteItem(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	responseBody := decodeDeleteData(writer, request)

	query := `DELETE FROM items WHERE user_id = $1 AND items.id = $2`

	_, err := datastore.DB.Exec(query, responseBody.UserId, responseBody.ItemId)
	if err != nil {
		err2.ErrorInternal(writer)
		panic(err)
	}
	writer.Write([]byte("Delete item SUCCESS"))
}

func decodeDeleteData(writer http.ResponseWriter, request *http.Request) dtos.DeleteItemData {
	decoder := json.NewDecoder(request.Body)
	var itemData dtos.DeleteItemData
	if err := decoder.Decode(&itemData); err != nil {
		err2.ErrorInternal(writer)
		panic(err)
	}
	return itemData
}
