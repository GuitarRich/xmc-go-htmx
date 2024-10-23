package sitecore

import (
	"fmt"
	"log"
	"os"

	"github.com/guitarrich/headless-go-htmx/model"
	"github.com/joho/godotenv"
)

func GetEnvVar(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}

func GetSafeString(field interface{}) string {
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

func IsEditMode(data model.SitecoreContext) bool {
	return data.EditMode == "edit"
}
