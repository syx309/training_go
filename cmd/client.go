package cmd

import (
	"log"
	"net/http"

	"github.com/syx309/training_go/cmd/helpers"
	"github.com/syx309/training_go/internal/services"

	"github.com/julienschmidt/httprouter"
)

func Start()  {
	helpers.InitDatabase()
	defer helpers.CloseDatabase()

	router := httprouter.New()
	router.POST("/register", services.RouteRegister)
	router.POST("/login", services.RouteLogin)

	router.POST("/users", helpers.BasicAuth(services.RouteUsers))
	router.POST("/user", helpers.BasicAuth(services.RouteUser))

	router.POST("/user/items", helpers.BasicAuth(services.RouteItems))
	router.POST("/user/item", helpers.BasicAuth(services.RouteGetItemByName))
	router.POST("/user/item/add", helpers.BasicAuth(services.RouteAddItem))
	router.POST("/user/item/update", helpers.BasicAuth(services.RouteUpdateItem))
	router.POST("/user/item/delete", helpers.BasicAuth(services.RouteDeleteItem))
	//router.POST("/ping", services.Ping)

	log.Fatal(http.ListenAndServe(":8080", router))
}
