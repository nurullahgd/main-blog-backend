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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Check if email already exists
	var existingUser models.User
	if err := database.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email already exists"})
	}

	// Check if username already exists
	if err := database.DB.Where("username = ?", input.Username).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Username already exists"})
	}

	hashedPassword, err := helpers.HashPassword(input.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	user := models.User{
		Name:      input.Name,
		Surname:   input.Surname,
		Username:  input.Username,
		Email:     input.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Password:  []byte(hashedPassword),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}
	c.Cookie(&fiber.Cookie{
		Name:     "user_token",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
	})

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})
}

func Login(c *fiber.Ctx) error {
	var input struct {
		Input    string `json:"input"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var user models.User
	if err := database.DB.Where("email = ? OR username = ?", input.Input, input.Input).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	if err := helpers.VerifyPassword(input.Password, string(user.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}
	c.Cookie(&fiber.Cookie{
		Name:     "user_token",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Login successful"})
}

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)

	// Convert to response format
	var response []models.UserResponse
	for _, user := range users {
		response = append(response, models.UserResponse{
			ID:           user.ID.String(),
			Name:         user.Name,
			Surname:      user.Surname,
			Username:     user.Username,
			Email:        user.Email,
			ProfileImage: user.ProfileImage,
			BlogCount:    user.BlogCount,
			CreatedAt:    user.CreatedAt,
			UpdatedAt:    user.UpdatedAt,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func GetUser(c *fiber.Ctx) error {
	var user models.User
	if err := database.DB.First(&user, "id = ?", c.Params("id")).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	response := models.UserResponse{
		ID:           user.ID.String(),
		Name:         user.Name,
		Surname:      user.Surname,
		Username:     user.Username,
		Email:        user.Email,
		ProfileImage: user.ProfileImage,
		BlogCount:    user.BlogCount,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func EditUser(c *fiber.Ctx) error {
	// Kullanıcı ID'sini cookie'den al
	userToken := c.Cookies("user_token")
	if userToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	// Token'ı doğrula ve kullanıcı ID'sini al
	userID, err := helpers.GetUserIDFromToken(userToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	// Kullanıcıyı bul
	var user models.User
	if err := database.DB.First(&user, "id = ?", userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// Form verilerini al
	name := c.FormValue("name")
	surname := c.FormValue("surname")

	// Sadece değişen alanları güncelle
	if name != "" {
		user.Name = name
	}
	if surname != "" {
		user.Surname = surname
	}

	// Kullanıcıyı güncelle
	if err := database.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update user",
		})
	}
	var userResponse models.UserResponse
	userResponse.ID = user.ID.String()
	userResponse.Name = user.Name
	userResponse.Surname = user.Surname
	userResponse.Username = user.Username
	userResponse.Email = user.Email
	userResponse.ProfileImage = user.ProfileImage
	userResponse.BlogCount = user.BlogCount

	return c.JSON(userResponse)
}

func UploadProfileImage(c *fiber.Ctx) error {
	userToken := c.Cookies("user_token")
	userID, err := helpers.GetUserIDFromToken(userToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Get user
	var user models.User
	if err := database.DB.First(&user, "id = ?", userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Get file from form
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No image file provided"})
	}

	// Validate file size (max 5MB)
	if file.Size > 5*1024*1024 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "File size too large. Maximum size is 5MB"})
	}

	// Upload to Cloudinary
	imageURL, err := utils.UploadToCloudinary(file, "profile_images")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Delete old image if exists
	if user.ProfileImage != "" {
		publicID := utils.GetPublicIDFromURL(user.ProfileImage)
		utils.DeleteFromCloudinary(publicID)
	}

	// Update user
	user.ProfileImage = imageURL
	if err := database.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save user profile"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Profile image updated successfully"})
}

func Logout(c *fiber.Ctx) error {
	// Clear the cookie
	c.Cookie(&fiber.Cookie{
		Name:     "user_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour), // Expire immediately
		HTTPOnly: true,
		Secure:   true,
		SameSite: "None",
		Path:     "/",
	})

	return c.JSON(fiber.Map{
		"message": "Logout successful",
	})
}
