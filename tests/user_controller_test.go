package tests

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/nurullahgd/main-blog-backend/controllers"
	"github.com/nurullahgd/main-blog-backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	// Create a new Fiber app
	app := fiber.New()

	// Setup the route
	app.Post("/api/register", controllers.Register)

	// Create test user data
	userData := models.UserCreate{
		Name:     "testuser",
		Surname:  "testuser",
		Username: "testuser",
		Email:    "test@example.com",
		Password: "testpassword",
	}

	// Convert user data to JSON
	jsonData, _ := json.Marshal(userData)

	// Create a test request
	req := httptest.NewRequest("POST", "/api/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	resp, _ := app.Test(req)

	// Assert the status code
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

	// Parse the response body
	var response map[string]interface{}
	err := json.NewDecoder(resp.Body).Decode(&response)

	// Assert that there was no error parsing the response
	assert.Nil(t, err)
}

func TestLogin(t *testing.T) {
	// Create a new Fiber app
	app := fiber.New()

	// Setup the route
	app.Post("/api/login", controllers.Login)

	// Create test login data
	loginData := map[string]interface{}{
		"email":    "test@example.com",
		"password": "testpassword",
	}

	// Convert login data to JSON
	jsonData, _ := json.Marshal(loginData)

	// Create a test request
	req := httptest.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	resp, _ := app.Test(req)

	// Assert the status code
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	// Parse the response body
	var response map[string]interface{}
	err := json.NewDecoder(resp.Body).Decode(&response)

	// Assert that there was no error parsing the response
	assert.Nil(t, err)

	// Assert that the response contains a token
	assert.Contains(t, response, "token")
}
