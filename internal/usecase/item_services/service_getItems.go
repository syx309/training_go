package item_services

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

func RouteItems(writer http.ResponseWriter, request *http.Request,_ httprouter.Params) {
	userData := decodeUserData(request)

	query := `SELECT id FROM users WHERE email = $1`
	row := datastore.DB.QueryRow(query, userData.Email)

	var user dtos.User
	err := row.Scan(&user.Id)
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

	var item dtos.Item
	rows, err := datastore.DB.Query("SELECT id, user_id, app_name, app_email, app_password FROM items WHERE $1", user.Id)
	if err != nil {
		fmt.Println("Query error")
		panic(err)
	}
	defer rows.Close()

	var items []dtos.Item
	for rows.Next() {
		err = rows.Scan(&item.Id, &item.User_id, &item.App_name, &item.App_email, &item.App_password)
		if err != nil {
			panic(err)
		}
		items = append(items, item)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	data, marshallErr := json.Marshal(items)

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
