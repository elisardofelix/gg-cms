package main

import (
	"encoding/json"
	"fmt"
	"gg-cms/Controllers"
	"gg-cms/DB"
	"gg-cms/DataRepos"
	"gg-cms/Middlewares"
	"gg-cms/Models"
	"gg-cms/Services"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

var (
	dbPostRepoTest = DataRepos.NewPostRepoTest()

	postServiceTest  = Services.NewPostService(dbPostRepoTest)

	postControllerTest  = Controllers.NewPostController(postServiceTest)
)



// This function is used to do setup before executing the test functions
func TestMain(m *testing.M) {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Run tests
	m.Run()

	//Teardown Temp DB for Testing
	client, err := DB.GetMongoDBClient()
	if err == nil {
		err = DB.DropMongoTestDB(client)
		if err != nil {
			fmt.Println("Error Drop DB ==> " + err.Error())
		}
	} else {
		fmt.Println("Error Drop DB ==> " + err.Error())
	}

	os.Exit(0)
}

func TestIntegrationCreateUserAuthorized(t *testing.T){
	// Create a response recorder
	w := httptest.NewRecorder()

	r := gin.Default()

	// Define the route similar to its definition in the routes file
	r.POST("/blog/post", Middlewares.AuthorizeJWT(), postControllerTest.SavePost)


	// Create a request to send to the above route
	postPayload := `{"title": "test1", "content": "Content ...", "permaLink": "test1"}`
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

	var data *Models.Post
	var postdata Models.Post
	err = json.Unmarshal([]byte(w.Body.String()), &data)
	if err != nil {
		fmt.Println("Error Data: " + err.Error())
		t.Fail()
		return
	}

	postdata, err = postServiceTest.Find(data.PermaLink)
	if err != nil || data.ID != postdata.ID {
		fmt.Println("Error DB Data: *" + err.Error())
		t.Fail()
	}
}

func TestIntegrationUpdateAuthorized(t *testing.T){

	postToCreate := Models.Post{
		"",
		"TitleTest",
		"Content ...",
		"temp-update-test",
		time.Now(),
		"efelix",
	}

	postCreated, _ := postServiceTest.Save(postToCreate)


	postCreated.Title = "TitleTest2"

	payload, err := json.Marshal(postCreated)

	fmt.Println(string(payload))

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

func TestIntegrationDeleteAuthorized(t *testing.T){

	postToCreate := Models.Post{
		"",
		"TitleDeleteTest",
		"Content ...",
		"temp-delete-test",
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

	var data *[]Models.Post
	err = json.Unmarshal([]byte(w.Body.String()), &data)

	// Test that the http status code is 200
	if w.Code != 200 || len(*data) <= 0 {
		fmt.Println("Error Request: not 200 : " + w.Body.String())
		t.Fail()
		return
	}
}

func TestUnauthorizeAccessMidleware(t *testing.T){
	// Create a response recorder
	w := httptest.NewRecorder()

	r := gin.Default()

	// Define the route similar to its definition in the routes file
	r.GET("/", Middlewares.AuthorizeJWT(), func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"enter" : "0",
		})
	})

	// Create a request to send to the above route
	req, err := http.NewRequest("GET", "/", nil )

	if err != nil {
		fmt.Println("Error Request: " + err.Error())
		return
	}
	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	var data *[]Models.Post
	err = json.Unmarshal([]byte(w.Body.String()), &data)

	// Test that the http status code is 401
	if w.Code != 401 {
		fmt.Println("Error Request: not 401")
		t.Fail()
		return
	}
}