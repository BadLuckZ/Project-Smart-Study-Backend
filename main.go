package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// createUserDatabase()

	app := fiber.New()

	// เปิดใช้งาน CORS Middleware ใน Go Fiber Framework เพื่อให้ API ของคุณรองรับการถูกเรียกข้ามโดเมน 
	// เช่น ทำให้ frontend: 3000 ทำงานกับ backend: 8080 ได้
	app.Use(cors.New())

	// Route that is protected by AuthMiddleware
	app.Get("/api/user/me", AuthMiddleware(), GetUserDataHandler)

	app.Listen(":8080")
}
