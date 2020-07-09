package main

import (
	"encoding/json"
	"fmt"
	"gg-cms/DTOs"
	"gg-cms/Middlewares"
	"gg-cms/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestIntegrationCreatePostAuthorized(t *testing.T){
	// Create a response recorder
	w := httptest.NewRecorder()

	r := gin.Default()

	// Define the route similar to its definition in the routes file
	r.POST("/blog/post", Middlewares.AuthorizeJWT(), postControllerTest.SavePost)


	// Create a request to send to the above route
	postPayload := `{"title": "test1", "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ornare arcu dui vivamus arcu felis bibendum. Amet dictum sit amet justo donec enim diam. Tellus mauris a diam maecenas sed enim. Tincidunt vitae semper quis lectus nulla. Diam in arcu cursus euismod quis. Odio ut enim blandit volutpat maecenas volutpat. Tempor nec feugiat nisl pretium fusce id velit ut tortor. Enim lobortis scelerisque fermentum dui faucibus in. Consectetur lorem donec massa sapien faucibus et molestie ac feugiat. Tellus mauris a diam maecenas sed enim. Enim sed faucibus turpis in eu mi bibendum neque. Lacus sed viverra tellus in hac. Egestas sed sed risus pretium quam vulputate dignissim suspendisse. Viverra nibh cras pulvinar mattis. Quis auctor elit sed vulputate mi sit amet. Amet nulla facilisi morbi tempus. Volutpat commodo sed egestas egestas fringilla. Amet nulla facilisi morbi tempus iaculis urna. Elementum integer enim neque volutpat ac tincidunt.", "permaLink": "test1"}`
	req, err := http.NewRequest("POST", "/blog/post", strings.NewReader(postPayload))
	req.Header = http.Header{}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(postPayload)))
	req.Header.Add("Authorization", "Bearer " + jwtService.GenerateToken("efelix", true))

	if err != nil {
		fmt.Println("Error Request: " + err.Error())
		return
	}
	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	// Test that the http status code is 401 Unauthorized
	if w.Code != 200{
		fmt.Println("Error Request: not 200")
		t.Fail()
		return
	}

	var data Models.Post
	var postData Models.Post
	err = json.Unmarshal([]byte(w.Body.String()), &data)
	if err != nil {
		fmt.Println("Error Data: " + err.Error())
		t.Fail()
		return
	}

	postData, err = postServiceTest.Find(data.PermaLink)
	if err != nil || data.ID != postData.ID {
		fmt.Println("Error DB Data: *" + err.Error())
		t.Fail()
	}
}

func TestIntegrationUpdatePostAuthorized(t *testing.T){

	postToCreate := Models.Post{
		"",
		"TitleTest",
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ornare arcu dui vivamus arcu felis bibendum. Amet dictum sit amet justo donec enim diam. Tellus mauris a diam maecenas sed enim. Tincidunt vitae semper quis lectus nulla. Diam in arcu cursus euismod quis. Odio ut enim blandit volutpat maecenas volutpat. Tempor nec feugiat nisl pretium fusce id velit ut tortor. Enim lobortis scelerisque fermentum dui faucibus in. Consectetur lorem donec massa sapien faucibus et molestie ac feugiat. Tellus mauris a diam maecenas sed enim. Enim sed faucibus turpis in eu mi bibendum neque. Lacus sed viverra tellus in hac. Egestas sed sed risus pretium quam vulputate dignissim suspendisse. Viverra nibh cras pulvinar mattis. Quis auctor elit sed vulputate mi sit amet. Amet nulla facilisi morbi tempus. Volutpat commodo sed egestas egestas fringilla. Amet nulla facilisi morbi tempus iaculis urna. Elementum integer enim neque volutpat ac tincidunt.",
		"temp-update-test",
		"Active",
		time.Now(),
		"efelix",
	}

	postCreated, _ := postServiceTest.Save(postToCreate)


	postCreated.Title = "TitleTest2"

	payload, err := json.Marshal(postCreated)

	updatePayload := string(payload)

	// Create a response recorder
	w := httptest.NewRecorder()

	r := gin.Default()

	// Define the route similar to its definition in the routes file
	r.PATCH("/blog/post", Middlewares.AuthorizeJWT(), postControllerTest.UpdatePost)


	// Create a request to send to the above route
	req, err := http.NewRequest("PATCH", "/blog/post", strings.NewReader(updatePayload))
	req.Header = http.Header{}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(updatePayload)))

	//Authorization
	req.Header.Add("Authorization", "Bearer " + jwtService.GenerateToken("efelix", true))

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

	var data *Models.Post
	var updatedData Models.Post
	err = json.Unmarshal([]byte(w.Body.String()), &data)
	if err != nil {
		fmt.Println("Error Data: " + err.Error())
		t.Fail()
		return
	}

	updatedData, err = postServiceTest.Find(data.PermaLink)
	if err != nil || data.Title != updatedData.Title || data.ID != updatedData.ID  {
		fmt.Println("Error DB Data: *" + err.Error())
		t.Fail()
	}
}

func TestIntegrationDeletePostAuthorized(t *testing.T){

	postToCreate := Models.Post{
		"",
		"TitleDeleteTest",
		"Content ...",
		"temp-delete-test",
		"Active",

		time.Now(),
		"efelix",
	}

	postCreated, _ := postServiceTest.Save(postToCreate)

	// Create a response recorder
	w := httptest.NewRecorder()

	r := gin.Default()

	// Define the route similar to its definition in the routes file
	r.DELETE("/blog/post/:id", Middlewares.AuthorizeJWT(), postControllerTest.DeletePost)


	// Create a request to send to the above route
	req, err := http.NewRequest("DELETE", "/blog/post/" + postCreated.ID, nil )
	req.Header = http.Header{}
	//Authorization
	req.Header.Add("Authorization", "Bearer " + jwtService.GenerateToken("efelix", true))

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

	var deleteData Models.Post

	deleteData, err = postServiceTest.Find(postCreated.PermaLink)
	if err.Error() != "mongo: no documents in result" || deleteData.ID != "" {
		fmt.Println("Error DB Data: *" + err.Error())
		t.Fail()
	}
}

func TestIntegrationGetAllPosts(t *testing.T){

	postToCreate := Models.Post{
		"",
		"TitleDeleteTest",
		"Content ...",
		"temp-delete-test",
		"Active",
		time.Now(),
		"efelix",
	}

	postServiceTest.Save(postToCreate)

	// Create a response recorder
	w := httptest.NewRecorder()

	r := gin.Default()

	// Define the route similar to its definition in the routes file
	r.GET("/blog", postControllerTest.FindAllActivePosts)


	// Create a request to send to the above route
	req, err := http.NewRequest("GET", "/blog", nil )

	if err != nil {
		fmt.Println("Error Request: " + err.Error())
		return
	}
	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	var data DTOs.Posts
	err = json.Unmarshal([]byte(w.Body.String()), &data)


	// Test that the http status code is 200
	if w.Code != 200 || len(data.Data) <= 0 || data.Total <= 0 {
		fmt.Println("Error Request: not 200 : " + w.Body.String())
		t.Fail()
		return
	}
}