package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurullahgd/main-blog-backend/controllers"
	"github.com/nurullahgd/main-blog-backend/middleware"
)

func SetupRoutes(app *fiber.App) {
	// User routes (user panel)
	userRoutes := app.Group("/api/users")
	userRoutes.Get("/", controllers.GetUsers)
	userRoutes.Get("/:id", controllers.GetUser)
	userRoutes.Post("/register", controllers.Register)
	userRoutes.Post("/login", controllers.Login)
	userRoutes.Post("/logout", controllers.Logout)

	// Protected user routes (user panel)
	protectedUserRoutes := userRoutes.Group("/", middleware.AuthMiddleware())
	protectedUserRoutes.Put("/edit", controllers.EditUser)
	protectedUserRoutes.Post("/profile-image", controllers.UploadProfileImage)

// Blog routes (blog panel)
blogRoutes := app.Group("/api/blogs")
blogRoutes.Get("/", controllers.GetBlogs)

// Daha spesifik route'lar önce gelmeli
// Protected blog routes (blog panel)
protectedBlogRoutes := blogRoutes.Group("/", middleware.AuthMiddleware())
protectedBlogRoutes.Get("/fetchBlogs", controllers.FetchMyBlogs)
protectedBlogRoutes.Post("/createBlog", controllers.CreateBlog)
protectedBlogRoutes.Post("/visibility/:id", controllers.ChangeVisibility)
protectedBlogRoutes.Post("/editBlog/:id", controllers.EditBlog)
protectedBlogRoutes.Delete("/:id", controllers.DeleteBlog)
protectedBlogRoutes.Post("/:id/main-image", controllers.UploadBlogImage)

// En sona bırak: ID ile yapılan get işlemi
blogRoutes.Get("/:id", controllers.GetBlog)

	// Admin routes (admin panel)
	adminRoutes := app.Group("/api/admin", middleware.AdminAuthMiddleware())
	adminRoutes.Get("/getUsers", controllers.GetUsers)
	adminRoutes.Delete("/blogDelete/:id", controllers.DeleteBlogFromAdmin)
	adminRoutes.Delete("/userDelete/:id", controllers.DeleteUserFromAdmin)

	adminRoutes.Get("/users", controllers.GetAdminUsers)
	adminRoutes.Post("/users", controllers.CreateAdminUser)
}
