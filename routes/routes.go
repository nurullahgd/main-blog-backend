package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurullahgd/main-blog-backend/controllers"
	"github.com/nurullahgd/main-blog-backend/middleware"
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
	userRoutes.Get("/", controllers.GetUsers)
	userRoutes.Get("/:id", controllers.GetUser)
	userRoutes.Post("/register", controllers.Register)
	userRoutes.Post("/login", controllers.Login)
	userRoutes.Post("/logout", controllers.Logout)

	// Protected user routes
	protectedUserRoutes := userRoutes.Group("/", middleware.AuthMiddleware())
	protectedUserRoutes.Put("/edit", controllers.EditUser)
	protectedUserRoutes.Post("/profile-image", controllers.UploadProfileImage)

	// Blog routes
	blogRoutes := app.Group("/api/blogs")
	blogRoutes.Get("/", controllers.GetBlogs)
	blogRoutes.Get("/:id", controllers.GetBlog)
	blogRoutes.Patch("/:id", controllers.EditBlog)
	blogRoutes.Delete("/:id", controllers.DeleteBlog)

	// Protected blog routes
	protectedBlogRoutes := blogRoutes.Group("/", middleware.AuthMiddleware())
	protectedBlogRoutes.Post("/createBlog", controllers.CreateBlog)
	protectedBlogRoutes.Post("/:id/main-image", controllers.UploadBlogImage)

	// Admin routes
	adminRoutes := app.Group("/api/admin", middleware.AdminAuthMiddleware())
	adminRoutes.Get("/users", controllers.GetAdminUsers)
	adminRoutes.Post("/users", controllers.CreateAdminUser)
	adminRoutes.Delete("/blogDelete/:id", controllers.DeleteBlogFromAdmin)
}
