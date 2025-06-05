package controllers

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nurullahgd/main-blog-backend/database"
	"github.com/nurullahgd/main-blog-backend/helpers"
	"github.com/nurullahgd/main-blog-backend/models"
	"github.com/nurullahgd/main-blog-backend/utils"
)

func Register(c *fiber.Ctx) error {
	var input models.UserCreate
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	hashedPassword, err := helpers.HashPassword(input.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	user := models.User{
		Name:     input.Name,
		Surname:  input.Surname,
		Username: input.Username,
		Email:    input.Email,
		Password: []byte(hashedPassword),
	}

	database.DB.Create(&user)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{"message": "User created successfully", "token": tokenString})
}

func Login(c *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// TODO: Add proper password hashing and verification
	if err := helpers.VerifyPassword(input.Password, string(user.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Create session token
	token := utils.GenerateToken() // Implement this function

	// Set cookie
	cookie := fiber.Cookie{
		Name:     "session",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "None",
		Path:     "/",
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Login successful",
		"user": models.UserResponse{
			ID:           user.ID,
			Name:         user.Name,
			Surname:      user.Surname,
			Email:        user.Email,
			ProfileImage: user.ProfileImage,
			BlogCount:    user.BlogCount,
			CreatedAt:    user.CreatedAt,
			UpdatedAt:    user.UpdatedAt,
		},
	})
}

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)

	// Convert to response format
	var response []models.UserResponse
	for _, user := range users {
		response = append(response, models.UserResponse{
			ID:           user.ID,
			Name:         user.Name,
			Surname:      user.Surname,
			Email:        user.Email,
			ProfileImage: user.ProfileImage,
			BlogCount:    user.BlogCount,
			CreatedAt:    user.CreatedAt,
			UpdatedAt:    user.UpdatedAt,
		})
	}

	return c.JSON(response)
}

func GetUser(c *fiber.Ctx) error {
	var user models.User
	if err := database.DB.First(&user, "id = ?", c.Params("id")).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	response := models.UserResponse{
		ID:           user.ID,
		Name:         user.Name,
		Surname:      user.Surname,
		Email:        user.Email,
		ProfileImage: user.ProfileImage,
		BlogCount:    user.BlogCount,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}

	return c.JSON(response)
}

func UploadProfileImage(c *fiber.Ctx) error {
	userID := c.Params("id")

	// Get user
	var user models.User
	if err := database.DB.First(&user, "id = ?", userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	// Get file from form
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "No image file provided"})
	}

	// Upload to Cloudinary
	imageURL, err := utils.UploadToCloudinary(file, "profile_images")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	// Delete old image if exists
	if user.ProfileImage != "" {
		publicID := utils.GetPublicIDFromURL(user.ProfileImage)
		utils.DeleteFromCloudinary(publicID)
	}

	// Update user
	user.ProfileImage = imageURL
	database.DB.Save(&user)

	response := models.UserResponse{
		ID:           user.ID,
		Name:         user.Name,
		Surname:      user.Surname,
		Email:        user.Email,
		ProfileImage: user.ProfileImage,
		BlogCount:    user.BlogCount,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}

	return c.JSON(response)
}
