package test

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/syx309/training_go/cmd/datastore"
	"github.com/syx309/training_go/internal/services"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func BaseTestItem(email string, appName string) *httptest.ResponseRecorder {
	datastore.InitDatabase()
	router := httprouter.New()
	router.POST("/user/item", services.RouteGetItemByName)

	r := strings.NewReader(fmt.Sprintf("{\"email\": \"%s\", \"appName\": \"%s\"}", email, appName))
	request, _ := http.NewRequest("POST", "/user/item", r)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	return response
}

func TestItemSuccess(t *testing.T) {
	response := BaseTestItem(emailSuccess, appNameSuccess)
	if status := response.Code; status != http.StatusOK {
		t.Log(status)
		t.Errorf("Error Occured")
	} // Expect Success
}

func TestItemFail(t *testing.T) {
	response := BaseTestItem(emailSuccess,appNameFail)
	if status := response.Code; status != http.StatusUnauthorized {
		t.Log(status)
		t.Errorf("Error Occured")
	} // Expect Fail
}
