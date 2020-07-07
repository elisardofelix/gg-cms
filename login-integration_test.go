package main

import (
	"fmt"
	"gg-cms/Models"
	"gg-cms/Services"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestIntegrationLoginUserAuthorized(t *testing.T){

	username := "user22"
	password := "112345678"

	userToCreate := Models.User{
		UserName: username,
		Password: Services.EncodePassword(username, password),
		Email: "blablabla2@gmail.com",
		IsAdmin: true,
		Status: "Active",
		CreatedDate: time.Now(),
		CreatedBy: "owner",
	}

	userServiceTest.Save(userToCreate)



	postPayload := fmt.Sprintf(`{ "userName" : "%s", "password":"%s" }`, username, password)
	// Create a response recorder
	w := httptest.NewRecorder()

	r := gin.Default()

	// Define the route similar to its definition in the routes file
	r.POST("/login/auth", loginControllerTest.Login)


	// Create a request to send to the above route
	req, err := http.NewRequest("POST", "/login/auth", strings.NewReader(postPayload))

	if err != nil {
		fmt.Println("Error Request: " + err.Error())
		return
	}
	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	// Test that the http status code is 200
	if w.Code != 200{
		fmt.Println("Error Request: not 200 : " + w.Body.String())
		t.Fail()
		return
	}
}
