package main

import (
	"context"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 1. อ่านค่า Authorization header
		authHeader := c.Get("Authorization")

		// 2. เช็กว่า header มีข้อมูล และขึ้นต้นด้วย Bearer หรือไม่
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing or invalid Authorization header (ต้องเป็น Bearer <token>)",
			})
		}

		// 3. แยกเอาเฉพาะ token (ตัด "Bearer " ออก)
		idToken := strings.TrimPrefix(authHeader, "Bearer ")

		// 4. เตรียมใช้ Firebase Admin SDK (เชื่อมต่อ Firebase)
		app := InitFirebaseApp()
		authClient, err := app.Auth(context.Background())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Firebase Auth client error",
			})
		}

		// 5. ตรวจสอบ (verify) idToken ว่าใช้ได้จริงไหม
		token, err := authClient.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired ID token",
			})
		}

		// 6. ถ้าสำเร็จ เก็บ uid และ email ใน context (c.Locals)
		c.Locals("uid", token.UID)
		c.Locals("email", token.Claims["email"])

		// 7. ไปต่อที่ handler ตัวถัดไป
		return c.Next()
	}
}