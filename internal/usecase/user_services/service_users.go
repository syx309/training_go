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

func RouteUsers(writer http.ResponseWriter, _ *http.Request,_ httprouter.Params) {
	//query select
	var user dtos.User
	rows, err := datastore.DB.Query("SELECT id, name, email, password FROM users")
	if err != nil {
		fmt.Println("Query error")
		panic(err)
	}
	defer rows.Close()

	var users []dtos.User
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	data, marshallErr := json.Marshal(users)

	if marshallErr != nil {
		err2.ErrorInternal(writer)
		panic(marshallErr)
	}

	_, _ = writer.Write(data)
}
