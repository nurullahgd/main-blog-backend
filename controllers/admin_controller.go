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
			ID:        admin.ID,
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
		ID:        adminUser.ID,
		Username:  adminUser.Username,
		Email:     adminUser.Email,
		Role:      adminUser.Role,
		CreatedAt: adminUser.CreatedAt,
		UpdatedAt: adminUser.UpdatedAt,
	}

	return c.Status(201).JSON(response)
}
