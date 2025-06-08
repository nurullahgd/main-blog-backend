package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurullahgd/main-blog-backend/database"
	"github.com/nurullahgd/main-blog-backend/models"
)

func GetAdminUsers(c *fiber.Ctx) error {
	var adminUsers []models.AdminUser
	database.DB.Find(&adminUsers)

	// Convert to response format
	var response []models.AdminUserResponse
	for _, admin := range adminUsers {
		response = append(response, models.AdminUserResponse{
			ID:        admin.ID.String(),
			Username:  admin.Username,
			Email:     admin.Email,
			Role:      admin.Role,
			CreatedAt: admin.CreatedAt,
			UpdatedAt: admin.UpdatedAt,
		})
	}

	return c.JSON(response)
}

func CreateAdminUser(c *fiber.Ctx) error {
	var input models.AdminUserCreate
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	adminUser := models.AdminUser{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password, // In a real application, you should hash the password
		Role:     input.Role,
	}

	database.DB.Create(&adminUser)

	response := models.AdminUserResponse{
		ID:        adminUser.ID.String(),
		Username:  adminUser.Username,
		Email:     adminUser.Email,
		Role:      adminUser.Role,
		CreatedAt: adminUser.CreatedAt,
		UpdatedAt: adminUser.UpdatedAt,
	}

	return c.Status(201).JSON(response)
}

func DeleteBlogFromAdmin(c *fiber.Ctx) error {
	blogID := c.Params("id")

	var blog models.Blog
	if err := database.DB.First(&blog, "id = ?", blogID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Blog not found"})
	}
	database.DB.Delete(&blog)

	return c.Status(200).JSON(fiber.Map{"message": "Blog deleted successfully"})
}

func DeleteUserFromAdmin(c *fiber.Ctx) error {
	userID := c.Params("id")

	var user models.User
	if err := database.DB.First(&user, "id = ?", userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	database.DB.Delete(&user)

	return c.Status(200).JSON(fiber.Map{"message": "User deleted successfully"})
}
