package Middlewares

import (
	"fmt"
	"gg-cms/Services"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Json Web Token (JWT) Middleware", func() {

	var (
		jwtService    Services.JWTService
	)

	BeforeSuite(func() {
		gin.SetMode(gin.TestMode)
		jwtService = Services.NewJWTService()
	})

	Describe("Validating signed token and data extraction", func() {

		Context("If the token is successfully validated", func() {

			It("should return Success code in response", func() {
				token := jwtService.GenerateToken("efelix", false)

				res, err := executeRequestMiddlewareTest(AuthorizeJWT(), token)


				Ω(err).Should(BeNil())
				Ω(res.Code).Should(Equal(200))
			})

			It("should get the username from the token", func() {
				token := jwtService.GenerateToken("efelix", true)
				res, err := executeRequestMiddlewareTest(AuthorizeJWTAdmin(), token)


				Ω(err).Should(BeNil())
				Ω(res.Code).Should(Equal(200))
			})

		})

		Context("If the token is unsuccessfully validated", func() {
			It("should return an Unauthorized code because invalid token", func() {
				token := "eyJhbGciI1NiIsInRVCJ9.eyJuYW1lIjo4cCI6CwiDQ2Ghvc3QifQ.VlSxaCxTsyR6XD-YTBOLAp58E"
				res, err := executeRequestMiddlewareTest(AuthorizeJWT(), token)

				// Error Should be nil
				Ω(err).Should(BeNil())
				// Test that the http status code is 401
				Ω(res.Code).Should(Equal(401))
			})

			It("should return an Unauthorized code becuse is not admin", func() {
				token := jwtService.GenerateToken("efelix", false)
				res, err := executeRequestMiddlewareTest(AuthorizeJWTAdmin(), token)

				// Error Should be nil
				Ω(err).Should(BeNil())
				// Test that the http status code is 401
				Ω(res.Code).Should(Equal(401))
			})

		})
	})
})

func executeRequestMiddlewareTest(handler gin.HandlerFunc, token string)  (*httptest.ResponseRecorder, error){
	// Create a response recorder
	w := httptest.NewRecorder()

	r := gin.Default()

	// Define the route similar to its definition in the routes file
	r.GET("/", handler, func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"enter" : "0",
		})
	})

	// Create a request to send to the above route
	req, err := http.NewRequest("GET", "/", nil )

	if err != nil {
		fmt.Println("Error Request: " + err.Error())
	}

	// Add Authorization header with token
	req.Header.Add("Authorization", "Bearer " + token)

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	fmt.Printf("Request Code: %d", w.Code)
	return w, err
}