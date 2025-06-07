package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nurullahgd/main-blog-backend/database"
	"github.com/nurullahgd/main-blog-backend/models"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get token from cookie
		token := c.Cookies("user_token")
		if token == "" {
			return c.Status(401).JSON(fiber.Map{
				"error": "Unauthorized - No token provided",
			})
		}

		// Parse and validate token
		claims := jwt.MapClaims{}
		parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !parsedToken.Valid {
			return c.Status(401).JSON(fiber.Map{
				"error": "Unauthorized - Invalid token",
			})
		}

		// Get user ID from token
		userID, ok := claims["user_id"].(string)
		if !ok {
			return c.Status(401).JSON(fiber.Map{
				"error": "Unauthorized - Invalid token claims",
			})
		}

		// Check if user exists
		var user models.User
		if err := database.DB.First(&user, "id = ?", userID).Error; err != nil {
			return c.Status(401).JSON(fiber.Map{
				"error": "Unauthorized - User not found",
			})
		}

		// Add user to context
		c.Locals("user", user)
		c.Locals("userID", userID)

		return c.Next()
	}
}

// OptionalAuthMiddleware is like AuthMiddleware but doesn't require authentication
func OptionalAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("user_token")
		if token == "" {
			return c.Next()
		}

		claims := jwt.MapClaims{}
		parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err == nil && parsedToken.Valid {
			if userID, ok := claims["user_id"].(string); ok {
				var user models.User
				if err := database.DB.First(&user, "id = ?", userID).Error; err == nil {
					c.Locals("user", user)
					c.Locals("userID", userID)
				}
			}
		}

		return c.Next()
	}
}

// CheckUserID checks if the requested user ID matches the authenticated user's ID
func CheckUserID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("userID")
		requestedID := c.Params("id")

		if userID != requestedID {
			return c.Status(403).JSON(fiber.Map{
				"error": "Forbidden - You can only access your own data",
			})
		}

		return c.Next()
	}
}

func AdminAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("admin_token")
		if token == "" {
			return c.Status(401).JSON(fiber.Map{
				"error": "Unauthorized - No token provided",
			})
		}

		claims := jwt.MapClaims{}
		parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err == nil && parsedToken.Valid {
			if userID, ok := claims["user_id"].(string); ok {
				var user models.AdminUser
				if err := database.DB.First(&user, "id = ?", userID).Error; err == nil {
					c.Locals("user", user)
					c.Locals("userID", userID)
				}
			}
		}

		return c.Next()
	}
}
