package handler

import (
	"github.com/guitarrich/headless-go-htmx/model"
	"github.com/guitarrich/headless-go-htmx/view/components"
)

func GetComponent(componentName string, component model.PlaceholderComponent) func(component model.PlaceholderComponent) (interface{}, error) {
	funcMap := map[string]interface{}{
		"Promo": components.Promo,
	}

	if funcMap[componentName] != nil {
		return nil // funcMap[componentName](component)
	}
	return nil
}
