package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"lab11/task10/go_service/handlers"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/health", handlers.GetHealth)
	router.GET("/items", handlers.GetItems)
	router.GET("/config", handlers.GetConfig)
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

func TestGetConfig(t *testing.T) {
	os.Setenv("APP_ENV", "test")
	os.Setenv("APP_VERSION", "1.0.0")
	os.Setenv("GO_PORT", "8080")

	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/config", nil)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("ожидался статус 200, получен %d", w.Code)
	}
}