package handlers

import (
	"os"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var items = []Item{
	{ID: 1, Name: "Item One"},
	{ID: 2, Name: "Item Two"},
	{ID: 3, Name: "Item Three"},
}

func GetItems(c *gin.Context) {
	c.JSON(200, items)
}

func GetHealth(c *gin.Context) {
	c.JSON(200, gin.H{"status": "healthy"})
}

func GetConfig(c *gin.Context) {
	c.JSON(200, gin.H{
		"app_env":     os.Getenv("APP_ENV"),
		"app_version": os.Getenv("APP_VERSION"),
		"port":        os.Getenv("GO_PORT"),
	})
}