package test

import (
	"github.com/julienschmidt/httprouter"
	"github.com/syx309/training_go/internal/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

func BaseTestUser(email string) *httptest.ResponseRecorder {
	//helpers.InitDatabase()
	router := httprouter.New()
	router.POST("/user", services.RouteUser)

	request, _ := http.NewRequest("POST", "/user", nil)
	request.Header.Set("email", email)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	return response
}

func TestUserSuccess(t *testing.T) {
	response := BaseTestUser(emailSuccess)
	if status := response.Code; status != http.StatusOK {
		t.Log(status)
		t.Errorf("Error Occured")
	} // Expect Success
}

func TestUserFail(t *testing.T) {
	response := BaseTestUser(emailFail)
	if status := response.Code; status != http.StatusUnauthorized {
		t.Log(status)
		t.Errorf("Error Occured")
	} // Expect Fail
}
