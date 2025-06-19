package main

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func InitFirebaseApp() *firebase.App {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Copy มาจาก Firebase SDK
	opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing firebase app: %v\n", err)
	}
	
	return app
}
