package sitecore

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnvVar(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}
