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

func RouteItems(writer http.ResponseWriter, request *http.Request,_ httprouter.Params) {
	userData := decodeUserData(request)

	query := `SELECT id FROM users WHERE email = $1`
	row := helpers.DB.QueryRow(query, userData.Email)

	var user dtos.User
	err := row.Scan(&user.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
			helpers.ErrorNotFound(writer)
			panic(err)
		} else {
			helpers.ErrorInternal(writer)
			panic(err)
		}
	}

	var item dtos.Item
	rows, err := helpers.DB.Query("SELECT id, user_id, app_name, app_email, app_password FROM items WHERE $1", user.Id)
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
		helpers.ErrorInternal(writer)
		panic(marshallErr)
	}

	_, _ = writer.Write(data)
}
