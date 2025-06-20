package main

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"github.com/gofiber/fiber/v2"
)

// ส่ง JSON ของ uid, email, username, photoUrl
func GetUserDataHandler(c *fiber.Ctx) error {
	// Get uid from c
	uid := c.Locals("uid").(string)

	// Get email from c
	email := c.Locals("email").(string)
	
	// ดึง displayName และ image จาก Firebase Admin SDK
	app := c.Locals("firebaseApp").(*firebase.App)
	client, err := app.Auth(context.Background())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "firebase error"})
	}

	user, err := client.GetUser(context.Background(), uid)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "cannot fetch user"})
	}

	return c.JSON(fiber.Map{
		"uid":		uid,
		"email":	email,
		"username": user.DisplayName,
		"photoUrl": user.PhotoURL,
	})
}