package Middlewares

import (
	"gg-cms/Services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// AuthorizeJWT validates the token from the http request, returning a 401 if it's not valid
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if(authHeader != ""){
			tokenString := authHeader[len(BEARER_SCHEMA):]

			token, err := Services.NewJWTService().ValidateToken(tokenString)

			if token.Valid {
				claims := token.Claims.(jwt.MapClaims)
				c.Set("jwtClaims", claims)
				/*
					//Just for example
					log.Println("Claims[Name]: ", claims["name"])
					log.Println("Claims[Admin]: ", claims["admin"])
					log.Println("Claims[Issuer]: ", claims["iss"])
					log.Println("Claims[IssuedAt]: ", claims["iat"])
					log.Println("Claims[ExpiresAt]: ", claims["exp"])
				 */
			} else {
				log.Println(err)
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)

		}

	}
}
