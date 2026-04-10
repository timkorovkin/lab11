package main

import (
	"lab11/task2/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.GET("/status", handlers.GetStatus)
	router.GET("/items", handlers.GetItems)

	router.Run(":8080")
}