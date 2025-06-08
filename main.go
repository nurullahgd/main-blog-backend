package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nurullahgd/main-blog-backend/database"
	"github.com/nurullahgd/main-blog-backend/models"
	"github.com/nurullahgd/main-blog-backend/routes"
	"github.com/nurullahgd/main-blog-backend/utils"
)

func main() {
	// Initialize database
	database.InitDB()

	// Auto Migrate the schema
	database.DB.AutoMigrate(&models.User{}, &models.Blog{}, &models.AdminUser{})

	// Initialize Cloudinary
	if err := utils.InitCloudinary(); err != nil {
		log.Fatal("Failed to initialize Cloudinary:", err)
	}

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		AppName:                 "Blog API v1.0",
		EnableTrustedProxyCheck: true,
		EnablePrintRoutes:       true,
	})

	// Add logger middleware
	app.Use(logger.New())

	// CORS configuration
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://your-frontend-domain.com,http://localhost:8000", // Frontend domaininizi buraya ekleyin
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,  // Important for cookies
		MaxAge:           43200, // 12 hours in seconds
	}))

	// Setup routes
	routes.SetupRoutes(app)

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Fatal(app.Listen(":" + port))
}
