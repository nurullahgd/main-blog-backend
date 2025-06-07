package controllers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	"github.com/nurullahgd/main-blog-backend/database"
	"github.com/nurullahgd/main-blog-backend/helpers"
	"github.com/nurullahgd/main-blog-backend/models"
	"github.com/nurullahgd/main-blog-backend/utils"
	"gorm.io/gorm"
)

func GetBlogs(c *fiber.Ctx) error {
	var blogs []models.Blog
	database.DB.Find(&blogs)

	// Convert to response format
	var response []models.BlogResponse
	for _, blog := range blogs {
		response = append(response, models.BlogResponse{
			ID:        blog.ID.String(),
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
		ID:        blog.ID.String(),
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
	userToken := c.Cookies("user_token")
	userID, err := helpers.GetUserIDFromToken(userToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Text verileri al
	title := c.FormValue("title")
	content := c.FormValue("content")
	visibilityStr := c.FormValue("visibility")
	slugInput := c.FormValue("slug")
	category := c.FormValue("category")
	summary := c.FormValue("summary")
	// Visibility değerini boolean'a çevir
	visibility := visibilityStr == "true" || visibilityStr == "1"

	// Slug oluştur
	var generatedSlug string
	if slugInput != "" {
		generatedSlug = slug.Make(slugInput)
	} else {
		generatedSlug = slug.Make(title)
	}

	// Slug benzersiz mi kontrol et
	uniqueSlug := generatedSlug
	counter := 1
	for {
		var existing models.Blog
		err := database.DB.Where("slug = ?", uniqueSlug).First(&existing).Error
		if err == gorm.ErrRecordNotFound {
			break
		}
		uniqueSlug = fmt.Sprintf("%s-%d", generatedSlug, counter)
		counter++
	}

	// Dosyayı al
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Image is required"})
	}

	// Cloudinary'e yükle
	imageURL, err := utils.UploadToCloudinary(file, "blogs")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Image upload failed"})
	}

	// DB'ye kaydet
	blog := models.Blog{
		Title:      title,
		Content:    content,
		MainImage:  imageURL,
		UserID:     userID,
		Visibility: visibility,
		Slug:       uniqueSlug,
		Category:   category,
		Summary:    summary,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := database.DB.Create(&blog).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create blog"})
	}

	// Blog sayısını arttır
	database.DB.Model(&models.User{}).Where("id = ?", userID).Update("blog_count", gorm.Expr("blog_count + 1"))

	// Yanıt
	response := models.BlogResponse{
		ID:         blog.ID.String(),
		Title:      blog.Title,
		Content:    blog.Content,
		Slug:       blog.Slug,
		MainImage:  blog.MainImage,
		UserID:     blog.UserID,
		Category:   blog.Category,
		Visibility: blog.Visibility,
		Summary:    blog.Summary,
		CreatedAt:  blog.CreatedAt,
		UpdatedAt:  blog.UpdatedAt,
	}

	return c.Status(fiber.StatusCreated).JSON(response)
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
		ID:        blog.ID.String(),
		Title:     blog.Title,
		Content:   blog.Content,
		MainImage: blog.MainImage,
		UserID:    blog.UserID,
		Summary:   blog.Summary,
		CreatedAt: blog.CreatedAt,
		UpdatedAt: blog.UpdatedAt,
	}

	return c.JSON(response)
}
