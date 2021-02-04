package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/syx309/training_go/cmd/datastore"
	"github.com/syx309/training_go/internal/auth"
	"github.com/syx309/training_go/internal/usecase/item_services"
	"github.com/syx309/training_go/internal/usecase/user_services"
	"log"
	"net/http"
)

func main(){
	datastore.InitDatabase()
	defer datastore.CloseDatabase()

	router := httprouter.New()
	router.POST("/register", user_services.RouteRegister)
	router.POST("/login", user_services.RouteLogin)

	router.POST("/users", auth.BasicAuth(user_services.RouteUsers))
	router.POST("/user", auth.BasicAuth(user_services.RouteUser))

	router.POST("/user/items", auth.BasicAuth(item_services.RouteItems))
	router.POST("/user/item", auth.BasicAuth(item_services.RouteGetItemByName))
	router.POST("/user/item/add", auth.BasicAuth(item_services.RouteAddItem))
	router.POST("/user/item/update", auth.BasicAuth(item_services.RouteUpdateItem))
	router.POST("/user/item/delete", auth.BasicAuth(item_services.RouteDeleteItem))
	//router.POST("/ping", usecase.Ping)

	log.Fatal(http.ListenAndServe(":8080", router))
}
