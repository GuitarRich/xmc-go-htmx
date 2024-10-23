package editing

import (
	"fmt"
	"net/http"

	"github.com/guitarrich/headless-go-htmx/sitecore/render"
	"github.com/labstack/echo/v4"
)

func (h *EditingRequestHandler) Config(c echo.Context) error {

	fmt.Println("editingRenderMiddlewareConfig")
	fmt.Printf("Request: %s\n", c.Request().URL.Path)

	// Validate allowed origins and editing secret
	if !enforceCors(c, EDITING_ALLOWED_ORIGINS) {
		message := fmt.Sprintf("Requests from origin %s not allowed", c.Request().Header.Get("Origin"))
		return c.HTML(http.StatusForbidden, message)
	}

	// Validate editing secret
	secret := c.Request().URL.Query().Get(QUERY_PARAM_EDITING_SECRET)
	editingSecret, err := getJssEditingSecret()
	if err != nil {
		return c.HTML(http.StatusForbidden, err.Error())
	}

	if secret != editingSecret {
		message := fmt.Sprintf("Invalid editing secret")
		return c.HTML(http.StatusForbidden, message)
	}

	c.Response().Header().Set("Content-Security-Policy", getSCHeader())

	packages := make(map[string]interface{})

	componentList := render.GetComponents()
	keys := make([]string, 0, len(componentList))

	for key := range componentList {
		keys = append(keys, key)
	}

	response := EditingMiddlewareConfig{
		Components:    keys,
		PagesEditMode: EditModeMetadata,
		Packages:      packages,
	}

	/*
		    responseJson := json.Marshal(response)

			c.Response().Header().Set("Content-Type", "application/json")
			_, err = c.Response().Write(responseJson)
			if err != nil {
				return err
			}
	*/
	return c.JSON(http.StatusOK, response)
}
