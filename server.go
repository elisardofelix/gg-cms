package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	server := gin.Default()
	port := os.Getenv("PORT")
	if port == ""{
		port = "8000"
	}
	initializeRoutes(server)
	server.Run(":" + port)
}
