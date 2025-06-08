package tests

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/nurullahgd/main-blog-backend/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetAllBlogs(t *testing.T) {
	// Create a new Fiber app
	app := fiber.New()

	// Setup the route
	app.Get("/api/blogs", controllers.GetBlogs)

	// Create a test request
	req := httptest.NewRequest("GET", "/api/blogs", nil)

	// Create a response recorder
	resp, _ := app.Test(req)

	// Assert the status code
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	// Parse the response body
	var response map[string]interface{}
	err := json.NewDecoder(resp.Body).Decode(&response)

	// Assert that there was no error parsing the response
	assert.Nil(t, err)

	// Add more specific assertions based on your expected response structure
}

func TestGetBlogByID(t *testing.T) {
	// Create a new Fiber app
	app := fiber.New()

	// Setup the route
	app.Get("/api/blogs/:id", controllers.GetBlog)

	// Create a test request with a sample ID
	req := httptest.NewRequest("GET", "/api/blogs/1", nil)

	// Create a response recorder
	resp, _ := app.Test(req)

	// Assert the status code
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	// Parse the response body
	var response map[string]interface{}
	err := json.NewDecoder(resp.Body).Decode(&response)

	// Assert that there was no error parsing the response
	assert.Nil(t, err)

	// Add more specific assertions based on your expected response structure
}
