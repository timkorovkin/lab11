package handlers

import "github.com/gin-gonic/gin"

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