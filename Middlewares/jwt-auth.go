package Middlewares

import (
	"gg-cms/Services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AuthorizeJWT() gin.HandlerFunc {
	return authorizeJWT(false)
}

func AuthorizeJWTAdmin() gin.HandlerFunc {
	return authorizeJWT(true)
}

// AuthorizeJWT validates the token from the http request, returning a 401 if it's not valid
func authorizeJWT(onlyAdmin bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if(authHeader != ""){
			tokenString := authHeader[len(BEARER_SCHEMA):]

			token, err := Services.NewJWTService().ValidateToken(tokenString)

			if token.Valid {
				claims := token.Claims.(jwt.MapClaims)
				c.Set("jwtClaims", claims)
				isAdmin := claims["admin"].(bool)
				if !isAdmin && onlyAdmin {
					log.Println(err)
					c.AbortWithStatus(http.StatusUnauthorized)
				}
			} else {
				log.Println(err)
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)

		}

	}
}
