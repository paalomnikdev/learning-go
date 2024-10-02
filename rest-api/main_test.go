package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

var testDB *sql.DB
var server *gin.Engine
var validCredentialsUser models.User
var authToken string
var eventId int64

func TestMain(m *testing.M) {
	db.InitDB(":memory:")
	defer db.DB.Close()

	gin.SetMode(gin.TestMode)
	server = gin.Default()
	routes.RegisterRoutes(server)

	testDB = db.DB

	validCredentialsUser = models.User{
		Email: "tester@example.com",
		Password: "Sup3rs3cr3t",
	}

	code := m.Run()

	os.Exit(code)
}


func TestSignup(t *testing.T) {
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
	userJson, _ := json.Marshal(validCredentialsUser)

	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(userJson))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var loginResponse map[string]any
	err := json.Unmarshal(w.Body.Bytes(), &loginResponse)
	assert.NoError(t, err)

	token, exists := loginResponse["token"].(string)
	assert.True(t, exists)
	assert.NotEmpty(t, token)
	authToken = token
}

func TestLoginFail(t *testing.T) {
	invalidCredentialsUser := models.User{
		Email: "tester@example.com",
		Password: "inv@l1d",
	}

	userJson, _ := json.Marshal(invalidCredentialsUser)

	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(userJson))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestCreateEvent(t *testing.T) {
	event := models.Event{
		Name: "Temp Test Event",
		Description: "Temp Test Event",
		Location: "Nowhere",
		DateTime: time.Now(),
	}

	eventJson, _ := json.Marshal(event)

	req, _ := http.NewRequest(http.MethodPost, "/events", bytes.NewBuffer(eventJson))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authToken)

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)


	var createEventResponse struct {
		Event models.Event `json:"event"`
		Message string `json:"message"`
	}

	err := json.Unmarshal(w.Body.Bytes(), &createEventResponse)
	assert.NoError(t, err)

	assert.Equal(t, createEventResponse.Message, "Event created.")
	assert.Equal(t, createEventResponse.Event.Name, event.Name)
	assert.Equal(t, createEventResponse.Event.Description, event.Description)
	assert.Equal(t, createEventResponse.Event.Location, event.Location)

	eventId = createEventResponse.Event.ID
}

func TestDeleteEventSuccess(t *testing.T) {
	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/events/%d", eventId), nil)
	req.Header.Set("Authorization", authToken)

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var deleteResponse map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &deleteResponse)
	assert.NoError(t, err)
	message, exists := deleteResponse["message"]
	assert.True(t, exists)
	assert.Equal(t, message, "Event deleted.")
}
