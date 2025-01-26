package middlewares

import (
	"product-service/internal/models"

	"github.com/gofiber/fiber/v2"
)

func IsSeller(c *fiber.Ctx) error {
	// Get the Authorization header
	user := c.Locals("user").(models.User)
	if !user.IsSeller {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "only seller is allow",
		})
	}
	return c.Next()
}
