package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()
	initializeRoutes(server)
	server.Run(":8000")
}
