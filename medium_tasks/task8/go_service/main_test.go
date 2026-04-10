package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"lab11/task8/go_service/handlers"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/health", handlers.GetHealth)
	router.GET("/items", handlers.GetItems)
	return router
}

func TestHealth(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("ожидался статус 200, получен %d", w.Code)
	}
}

func TestGetItems(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items", nil)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("ожидался статус 200, получен %d", w.Code)
	}
}