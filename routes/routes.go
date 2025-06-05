package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurullahgd/main-blog-backend/controllers"
)

func SetupRoutes(app *fiber.App) {
	// CORS middleware
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Method() == "OPTIONS" {
			return c.SendStatus(204)
		}
		return c.Next()
	})

	// User routes
	userRoutes := app.Group("/api/users")
	userRoutes.Get("/getUsers", controllers.GetUsers)
	userRoutes.Get("/getUser/:id", controllers.GetUser)
	userRoutes.Post("/register", controllers.Register)
	userRoutes.Post("/login", controllers.Login)

	// Blog routes
	blogRoutes := app.Group("/api/blogs")
	blogRoutes.Get("/", controllers.GetBlogs)
	blogRoutes.Get("/:id", controllers.GetBlog)
	blogRoutes.Post("/", controllers.CreateBlog)
	blogRoutes.Post("/:id/main-image", controllers.UploadBlogImage)

	// Admin routes
	adminRoutes := app.Group("/api/admin")
	adminRoutes.Get("/users", controllers.GetAdminUsers)
	adminRoutes.Post("/users", controllers.CreateAdminUser)
}
