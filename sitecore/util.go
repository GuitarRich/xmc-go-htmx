package sitecore

import (
	"fmt"
	"os"

	"github.com/guitarrich/headless-go-htmx/model"
	"github.com/joho/godotenv"
)

func GetEnvVar(key string) string {
	godotenv.Load()
	result := os.Getenv(key)
	return result
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
