package services

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Ping(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	fmt.Println("PING!!!")
	writer.Write([]byte("PING!!!"))
}
