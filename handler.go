package main

import "github.com/gofiber/fiber/v2"

func GetUserMeHandler(c *fiber.Ctx) error {
	// Get uid from c
	uid := c.Locals("uid")

	// Get email from c
	email := c.Locals("email")
	
	return c.JSON(fiber.Map{
		"uid":   uid,
		"email": email,
	})
}