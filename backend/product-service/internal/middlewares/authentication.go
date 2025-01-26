package middlewares

import (
	"product-service/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func Authentication(c *fiber.Ctx) error {
	// Get the Authorization header
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	user, err := utils.DecodedJWT(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	c.Locals("user", user)
	return c.Next()
}
