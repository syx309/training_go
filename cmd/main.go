package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/syx309/training_go/cmd/datastore"
	"github.com/syx309/training_go/internal/auth"
	"github.com/syx309/training_go/internal/services"
	"log"
	"net/http"
)

func main(){
	datastore.InitDatabase()
	defer datastore.CloseDatabase()

	router := httprouter.New()
	router.POST("/register", services.RouteRegister)
	router.POST("/login", services.RouteLogin)

	router.POST("/users", auth.BasicAuth(services.RouteUsers))
	router.POST("/user", auth.BasicAuth(services.RouteUser))

	router.POST("/user/items", auth.BasicAuth(services.RouteItems))
	router.POST("/user/item", auth.BasicAuth(services.RouteGetItemByName))
	router.POST("/user/item/add", auth.BasicAuth(services.RouteAddItem))
	router.POST("/user/item/update", auth.BasicAuth(services.RouteUpdateItem))
	router.POST("/user/item/delete", auth.BasicAuth(services.RouteDeleteItem))
	//router.POST("/ping", services.Ping)

	log.Fatal(http.ListenAndServe(":8080", router))
}
