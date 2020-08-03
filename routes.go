// routes.go

package main

import (
	"gg-cms/Controllers"
	"gg-cms/DataRepos"
	"gg-cms/Middlewares"
	"gg-cms/Services"
	"github.com/gin-gonic/gin"
)

var (
	dbPostRepo = DataRepos.NewPostRepo()
	dbUserRepo = DataRepos.NewUserRepo()

	postService  = Services.NewPostService(dbPostRepo)
	userService  = Services.NewUserService(dbUserRepo)
	loginService = Services.NewLoginService(dbUserRepo)
	jwtService   = Services.NewJWTService()

	postController  = Controllers.NewPostController(postService)
	userController  = Controllers.NewUserController(userService)
	loginController = Controllers.NewLoginController(loginService, jwtService)
)

func CORS() gin.HandlerFunc {
	// TO allow CORS
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, skip, limit")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func initializeRoutes(router *gin.Engine) {

	router.Use(CORS())


	// Handle the index route
	router.GET("/", func(ctx *gin.Context){
		ctx.JSON(200, gin.H {
			"menssage" : "OK Im Here, Third change!!",
		})
	})

	// Group blog related routes together
	articleRoutes := router.Group("/blog")
	{
		articleRoutes.GET("/", postController.FindAllActivePosts)
		articleRoutes.GET("/post/:permalink", postController.FindPost)
		articleRoutes.POST("/post", Middlewares.AuthorizeJWT(), postController.SavePost)
		articleRoutes.PATCH("/post", Middlewares.AuthorizeJWT(), postController.UpdatePost)
		articleRoutes.DELETE("/post/:id", Middlewares.AuthorizeJWT(), postController.DeletePost)
	}

	userRoutes := router.Group("/user", Middlewares.AuthorizeJWTAdmin())
	{
		userRoutes.GET("/", userController.FindAllUsers)
		userRoutes.GET("/:username", userController.FindUser)
		userRoutes.POST("/", userController.SaveUser)
		userRoutes.PATCH("/", userController.UpdateUser)
		userRoutes.DELETE("/:id", userController.DeleteUser)
	}


	// Group user related routes together
	loginRoutes := router.Group("/login")
	{
		loginRoutes.POST("/auth", loginController.Login)
		loginRoutes.POST("/register", func(ctx *gin.Context){
			ctx.JSON(500, gin.H {
				"error" : "Register not implemented yet",
			})
		})

	}


}
