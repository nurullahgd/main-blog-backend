package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurullahgd/main-blog-backend/database"
	"github.com/nurullahgd/main-blog-backend/models"
	"github.com/nurullahgd/main-blog-backend/utils"
)

func GetBlogs(c *fiber.Ctx) error {
	var blogs []models.Blog
	database.DB.Find(&blogs)

	// Convert to response format
	var response []models.BlogResponse
	for _, blog := range blogs {
		response = append(response, models.BlogResponse{
			ID:        blog.ID,
			Title:     blog.Title,
			Content:   blog.Content,
			MainImage: blog.MainImage,
			UserID:    blog.UserID,
			CreatedAt: blog.CreatedAt,
			UpdatedAt: blog.UpdatedAt,
		})
	}

	return c.JSON(response)
}

func GetBlog(c *fiber.Ctx) error {
	var blog models.Blog
	if err := database.DB.First(&blog, "id = ?", c.Params("id")).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Blog not found"})
	}

	response := models.BlogResponse{
		ID:        blog.ID,
		Title:     blog.Title,
		Content:   blog.Content,
		MainImage: blog.MainImage,
		UserID:    blog.UserID,
		CreatedAt: blog.CreatedAt,
		UpdatedAt: blog.UpdatedAt,
	}

	return c.JSON(response)
}

func CreateBlog(c *fiber.Ctx) error {
	var input models.BlogCreate
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	blog := models.Blog{
		Title:   input.Title,
		Content: input.Content,
		UserID:  input.UserID,
	}

	database.DB.Create(&blog)

	response := models.BlogResponse{
		ID:        blog.ID,
		Title:     blog.Title,
		Content:   blog.Content,
		MainImage: blog.MainImage,
		UserID:    blog.UserID,
		CreatedAt: blog.CreatedAt,
		UpdatedAt: blog.UpdatedAt,
	}

	return c.Status(201).JSON(response)
}

func UploadBlogImage(c *fiber.Ctx) error {
	blogID := c.Params("id")

	// Get blog
	var blog models.Blog
	if err := database.DB.First(&blog, "id = ?", blogID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Blog not found"})
	}

	// Get file from form
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "No image file provided"})
	}

	// Upload to Cloudinary
	imageURL, err := utils.UploadToCloudinary(file, "blog_images")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	// Delete old image if exists
	if blog.MainImage != "" {
		publicID := utils.GetPublicIDFromURL(blog.MainImage)
		utils.DeleteFromCloudinary(publicID)
	}

	// Update blog
	blog.MainImage = imageURL
	database.DB.Save(&blog)

	response := models.BlogResponse{
		ID:        blog.ID,
		Title:     blog.Title,
		Content:   blog.Content,
		MainImage: blog.MainImage,
		UserID:    blog.UserID,
		CreatedAt: blog.CreatedAt,
		UpdatedAt: blog.UpdatedAt,
	}

	return c.JSON(response)
}
