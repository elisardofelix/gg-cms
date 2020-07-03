package Services
import (
	"github.com/dgrijalva/jwt-go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

)
var _ = Describe("Json Web Token (JWT) Service", func() {

	var (
		jwtService    JWTService

	)

	BeforeSuite(func() {
		jwtService = NewJWTService()
	})

	Describe("Validating token validation and data extraction", func() {

		Context("If the token is successfully validated", func() {

			It("should return true in the token validation function", func() {
				token := jwtService.GenerateToken("efelix", true)
				_, err := jwtService.ValidateToken(token)

				Ω(err).Should(BeNil())
			})

			It("should get the username from the token", func() {
				username := "efelix"

				token := jwtService.GenerateToken(username, true)
				jwttoken, _ := jwtService.ValidateToken(token)
				claims := jwttoken.Claims.(jwt.MapClaims)

				Ω(claims["name"]).Should(Equal(username))
			})



		})

		Context("If the token is unsuccessfully validated", func() {

			It("should return an error", func() {
				token := "badtoken"
				_, err := jwtService.ValidateToken(token)

				Ω(err).ShouldNot(BeNil())
			})

		})
	})
})
