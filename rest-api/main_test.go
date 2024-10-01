package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

var testDB *sql.DB
var validCredentialsUser models.User
var invalidCredentialsUser models.User

func TestMain(m *testing.M) {
	db.InitDB(":memory:")
	defer db.DB.Close()

	testDB = db.DB

	validCredentialsUser = models.User{
		Email: "tester@example.com",
		Password: "Sup3rs3cr3t",
	}

	invalidCredentialsUser = models.User{
		Email: "tester@example.com",
		Password: "inv@l1d",
	}

	code := m.Run()

	os.Exit(code)
}


func TestSignup(t *testing.T) {
	gin.SetMode(gin.TestMode)
	server := gin.Default()

	server.POST("/signup", routes.Signup)

	userJson, _ := json.Marshal(validCredentialsUser)

	req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(userJson))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	var count int
	err := db.DB.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", validCredentialsUser.Email).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)
}

func TestLoginSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)
	server := gin.Default()

	server.POST("/login", routes.Login)

	userJson, _ := json.Marshal(validCredentialsUser)

	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(userJson))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var loginResponse map[string]any
	err := json.Unmarshal(w.Body.Bytes(), &loginResponse)
	assert.NoError(t, err)

	token, exists := loginResponse["token"]
	assert.True(t, exists)
	assert.NotEmpty(t, token)
}

func TestLoginFail(t *testing.T) {
	gin.SetMode(gin.TestMode)
	server := gin.Default()

	server.POST("/login", routes.Login)

	userJson, _ := json.Marshal(invalidCredentialsUser)

	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(userJson))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
