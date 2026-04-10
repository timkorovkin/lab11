package main

import (
	"lab11/task10/go_service/handlers"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.GET("/health", handlers.GetHealth)
	router.GET("/items", handlers.GetItems)
	router.GET("/config", handlers.GetConfig)

	port := os.Getenv("GO_PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}