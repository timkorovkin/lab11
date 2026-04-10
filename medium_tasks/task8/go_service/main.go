package main

import (
	"lab11/task8/go_service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.GET("/health", handlers.GetHealth)
	router.GET("/items", handlers.GetItems)

	router.Run(":8080")
}