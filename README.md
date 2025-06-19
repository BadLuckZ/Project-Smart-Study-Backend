# My Adaptive Question Generator & Problem Solver - Backend

โปรเจกต์ Backend สำหรับระบบสร้างและฝึกฝนชุดคำถามอัตโนมัติด้วย AI  
พัฒนาด้วย Go โดยใช้ Go Fiber เป็น Web Framework

## Tech Stack

- Go 1.22+
- Fiber (github.com/gofiber/fiber/v2)
- (เตรียมต่อยอด: GORM, PostgreSQL, Firebase Admin SDK, Google Gemini API, unidoc/unipdf)

## Getting Started

1. Clone Repository

   ```bash
   git clone https://github.com/BadLuckZ/Project-Smart-Study-Backend.git
   cd Project-Smart-Study-Backend
   ```

2. ติดตั้ง Dependency

   ```bash
   go mod download
   ```

3. Run Server
   ```bash
   go run main.go
   ```
   เซิร์ฟเวอร์จะเริ่มที่ `localhost:3000` (ค่าเริ่มต้น)

---

## Workday Progress

### Day 1

- สร้าง Go Project ด้วย `go mod init`
- ติดตั้ง Go Fiber (`github.com/gofiber/fiber/v2`)
- สร้างไฟล์ `main.go` สำหรับรัน Fiber Server
- ทดสอบรันเซิร์ฟเวอร์ (`go run main.go`) และเข้าถึง root endpoint ได้
- สร้างไฟล์ `.gitignore` และ `README.md` สำหรับ Go โปรเจกต์
- Push โค้ดเริ่มต้นขึ้น GitHub

### Day 2

- ติดตั้ง GORM (`gorm.io/gorm`, `gorm.io/driver/postgres`) สำหรับ ORM
- ติดตั้ง `github.com/joho/godotenv`
- เขียนโค้ดเชื่อมต่อกับฐานข้อมูล PostgreSQL (อ่านค่าจาก .env)
- ทดสอบการเชื่อมต่อฐานข้อมูล หากเชื่อมต่อไม่ได้ต้อง log
- สร้าง Model `User` (ID, Email, Name, CreatedAt, UpdatedAt)
- ใช้ GORM AutoMigrate สำหรับตาราง users
- สร้างไฟล์ `.env` สำหรับตัวแปรที่ใช้เข้าถึงฐานข้อมูล
- Push โค้ดทั้งหมดขึ้น GitHub

### Day 3

- ไม่มีส่วนที่แก้ไข Backend

### Day 4

- ติดตั้ง และตั้งค่า **Firebase Admin SDK for Go** ผ่าน https://firebase.google.com/docs/admin/setup#go
- สร้างไฟล์แยกสำหรับ initialize Firebase Admin (`firebase_admin.go`)
- เขียน **Auth Middleware** (อ่าน header `Authorization: Bearer <token>`, verify userToken, inject uid และ email)
- สร้าง **Protected Endpoint** `/api/user/me` (GET) ที่คืน uid/email เฉพาะผู้ที่ Auth แล้ว (`middleware.go`)
- เพิ่ม `GOOGLE_APPLICATION_CREDENTIALS` เข้าไปใน `.env` และเพิ่ม `env.md` เพื่อบอกว่าต้อง Set อะไรบ้าง

---

## หมายเหตุ

- พัฒนาเพื่อใช้ร่วมกับ Frontend: [Project-Smart-Study-Frontend](https://github.com/BadLuckZ/Project-Smart-Study-Frontend)
