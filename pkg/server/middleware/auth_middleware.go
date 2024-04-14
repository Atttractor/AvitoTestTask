package middleware

import "github.com/gofiber/fiber/v2"

func IsAdmin(c *fiber.Ctx) error {
	token := c.Get("Token")
	if token != "admin_token" {
		return fiber.NewError(fiber.StatusUnauthorized)
	}
	return c.Next()
}

func IsUser(c *fiber.Ctx) error {
	token := c.Get("Token")
	if token == "user_token" || token != "admin_token" {
		return c.Next()
	}
	return fiber.NewError(fiber.StatusUnauthorized)
}