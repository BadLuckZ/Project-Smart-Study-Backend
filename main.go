package main

import "github.com/gofiber/fiber/v2"

func main() {
	// createUserDatabase()

	app := fiber.New()

	// Route that is protected by AuthMiddleware
	app.Get("/api/user/me", AuthMiddleware(), GetUserMeHandler)

	app.Listen(":8080")
}
