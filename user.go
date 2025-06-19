package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func createUserDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Load .env Error!")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"))

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection Successful", db)

	db.AutoMigrate(&User{}) // Automatically create User Table...

	fmt.Println("Create/Update User Table Successful")
}