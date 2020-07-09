package main

import (
	"encoding/json"
	"fmt"
	"gg-cms/DTOs"
	"gg-cms/Middlewares"
	"gg-cms/Models"
	"gg-cms/Services"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)



func TestIntegrationCreateUserAuthorized(t *testing.T){
	// Create a response recorder
	w := httptest.NewRecorder()

	r := gin.Default()

	// Define the route similar to its definition in the routes file
	r.POST("/user", Middlewares.AuthorizeJWTAdmin(), userControllerTest.SaveUser)


	userToCreate := DTOs.Registration{
		UserName: "user",
		Password: "123456",
		RePasword: "123456",
		Email: "elisardofelix@gmail.com",
		IsAdmin: true,
		Status: "Active",
	}

	payload, err := json.Marshal(userToCreate)
	postPayload := string(payload)

	// Create a request to send to the above route
	req, err := http.NewRequest("POST", "/user", strings.NewReader(postPayload))
	req.Header = http.Header{}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(postPayload)))
	req.Header.Add("Authorization", "Bearer " + jwtService.GenerateToken("user", true))

	if err != nil {
		fmt.Println("Error Request: " + err.Error())
		return
	}
	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	// Test that the http status code is 401 Unauthorized
	if w.Code != 200{
		fmt.Println("Error Request: not 200 : " + w.Body.String())
		t.Fail()
		return
	}

	var data DTOs.User
	err = json.Unmarshal([]byte(w.Body.String()), &data)
	if err != nil {
		fmt.Println("Error Data: " + err.Error())
		t.Fail()
		return
	}

	var postData DTOs.User
	postData, err = userServiceTest.Find(data.UserName)
	if err != nil || data.ID != postData.ID {
		fmt.Println("Error DB Data: *" + err.Error())
		t.Fail()
	}
}

func TestIntegrationUpdateUserAuthorized(t *testing.T){

	username := "user2"
	password := "112345678"

	userToCreate := Models.User{
		UserName: username,
		Password: Services.EncodePassword(username, password),
		Email: "blablabla@gmail.com",
		IsAdmin: true,
		Status: "Active",
		CreatedDate: time.Now(),
		CreatedBy: "owner",
	}

	userCreated, _ := userServiceTest.Save(userToCreate)

	userCreated.Status = "Inactive"

	payload, err := json.Marshal(userCreated)

	updatePayload := string(payload)

	// Create a response recorder
	w := httptest.NewRecorder()

	r := gin.Default()

	// Define the route similar to its definition in the routes file
	r.PATCH("/user", Middlewares.AuthorizeJWTAdmin(), userControllerTest.UpdateUser)


	// Create a request to send to the above route
	req, err := http.NewRequest("PATCH", "/user", strings.NewReader(updatePayload))
	req.Header = http.Header{}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(updatePayload)))
	req.Header.Add("Authorization", "Bearer " + jwtService.GenerateToken("user", true))

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

	var data *DTOs.User
	var updatedData DTOs.User
	err = json.Unmarshal([]byte(w.Body.String()), &data)
	if err != nil {
		fmt.Println("Error Data: " + err.Error())
		t.Fail()
		return
	}

	updatedData, err = userServiceTest.Find(data.UserName)
	if err != nil || data.Status != updatedData.Status  {
		fmt.Println("Error DB Data: *" + err.Error())
		t.Fail()
	}
}


func TestIntegrationDeleteUserAuthorized(t *testing.T){

	username := "user3"
	password := "1123456789"

	userToCreate := Models.User{
		UserName: username,
		Password: Services.EncodePassword(username, password),
		Email: "blablablabla@outlook.com",
		IsAdmin: true,
		Status: "Active",
		CreatedDate: time.Now(),
		CreatedBy: "owner",
	}
	userCreated, _ := userServiceTest.Save(userToCreate)

	// Create a response recorder
	w := httptest.NewRecorder()

	r := gin.Default()

	// Define the route similar to its definition in the routes file
	r.DELETE("/user/:id", Middlewares.AuthorizeJWTAdmin(), userControllerTest.DeleteUser)


	// Create a request to send to the above route
	req, err := http.NewRequest("DELETE", "/user/" + userCreated.ID, nil )
	req.Header = http.Header{}
	//Authorization
	req.Header.Add("Authorization", "Bearer " + jwtService.GenerateToken("user", true))

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

	var deleteData DTOs.User

	deleteData, err = userServiceTest.Find(userCreated.UserName)
	if err.Error() != "mongo: no documents in result" || deleteData.ID != "" {
		fmt.Println("Error DB Data: *" + err.Error())
		t.Fail()
	}
}


func TestIntegrationGetAllUsers(t *testing.T){

	username := "user3"
	password := "1123456789"

	userToCreate := Models.User{
		UserName: username,
		Password: Services.EncodePassword(username, password),
		Email: "blablablabla@outlook.com",
		IsAdmin: true,
		Status: "Active",
		CreatedDate: time.Now(),
		CreatedBy: "owner",
	}

	userServiceTest.Save(userToCreate)

	// Create a response recorder
	w := httptest.NewRecorder()

	r := gin.Default()

	// Define the route similar to its definition in the routes file
	r.GET("/user", Middlewares.AuthorizeJWTAdmin(), userControllerTest.FindAllUsers)


	// Create a request to send to the above route
	req, err := http.NewRequest("GET", "/user", nil )
	req.Header = http.Header{}
	//Authorization
	req.Header.Add("Authorization", "Bearer " + jwtService.GenerateToken("user", true))

	if err != nil {
		fmt.Println("Error Request: " + err.Error())
		return
	}
	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	var data DTOs.Users
	err = json.Unmarshal([]byte(w.Body.String()), &data)

	// Test that the http status code is 200
	if w.Code != 200 || len(data.Data) <= 0 || data.Total <= 0 {
		fmt.Println("Error Request: not 200 : " + w.Body.String())
		t.Fail()
		return
	}
}
