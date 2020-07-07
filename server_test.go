package main

import (
	"fmt"
	"gg-cms/Controllers"
	"gg-cms/DB"
	"gg-cms/DataRepos"
	"gg-cms/Services"
	"github.com/gin-gonic/gin"
	"os"
	"testing"
)

var (
	dbPostRepoTest = DataRepos.NewPostRepoTest()
	dbUserRepoTest = DataRepos.NewUserRepoTest()

	postServiceTest  = Services.NewPostService(dbPostRepoTest)
	userServiceTest  = Services.NewUserService(dbUserRepoTest)
	loginServiceTest = Services.NewLoginService(dbUserRepoTest)

	postControllerTest  = Controllers.NewPostController(postServiceTest)
	userControllerTest  = Controllers.NewUserController(userServiceTest)
	loginControllerTest = Controllers.NewLoginController(loginServiceTest, jwtService)
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

