package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"neon_go/config"
	"neon_go/handlers"
)

func main() {
	// Set up database connection
	config.ConnectDB()

	// Define routes
	router := gin.Default()
	router.GET("/data/", handlers.HandleGetAll)
	router.GET("/data/:id", handlers.HandleGet)
	router.POST("/data", handlers.HandlePost)
	router.PUT("/data/:id", handlers.HandlePut)
	router.DELETE("/data/:id", handlers.HandleDelete)

	// Start the server
	err := router.Run(":8080")
	if err != nil {
		fmt.Print(err)
		return
	}
}
