package sitecore

import (
	"fmt"
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

func getSafeString(field interface{}) string {
	if field == nil {
		return ""
	}
	return fmt.Sprintf("%s", field)
}

func AddIfNotEmpty(name string, value string) string {
	if value != "" {
		return fmt.Sprintf(" %s=\"%s\"", name, value)
	}
	return ""
}
