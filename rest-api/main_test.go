package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)


func TestSignupAndLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db.InitDB(":memory:")
	defer db.DB.Close()

	server := gin.Default()
	server.POST("/signup", routes.Signup)
	server.POST("/login", routes.Login)

	user := models.User{
		Email: "tester@example.com",
		Password: "supers3cr3t",
	}

	userJson, _ := json.Marshal(user)

	// test signup
	req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(userJson))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	var count int
	err := db.DB.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", user.Email).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	//test success login
	req, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(userJson))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()

	server.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var loginResponse map[string]any
	err = json.Unmarshal(w.Body.Bytes(), &loginResponse)
	assert.NoError(t, err)

	token, exists := loginResponse["token"]
	assert.True(t, exists)
	assert.NotEmpty(t, token)

	//test failed login
	user.Password = "invalidPassword"
	invalidUserJson, _ := json.Marshal(user)
	req, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(invalidUserJson))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	server.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
