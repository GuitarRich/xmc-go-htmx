package editing

import (
	"fmt"
	"net/http"

	"github.com/guitarrich/headless-go-htmx/view/components"
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
	packages["@sitecore/byoc"] = "0.2.15"
	packages["@sitecore/components"] = "1.1.10"
	packages["@sitecore-cloudsdk/core"] = "0.3.1"
	packages["@sitecore-cloudsdk/events"] = "0.3.1"
	packages["@sitecore-cloudsdk/personalize"] = "0.3.1"
	packages["@sitecore-cloudsdk/utils"] = "0.3.1"
	packages["@sitecore-feaas/clientside"] = "0.5.18"
	packages["@sitecore-jss/sitecore-jss"] = "22.1.3"
	packages["@sitecore-jss/sitecore-jss-cli"] = "22.1.3"
	packages["@sitecore-jss/sitecore-jss-dev-tools"] = "22.1.3"
	packages["@sitecore-jss/sitecore-jss-nextjs"] = "22.1.3"
	packages["@sitecore-jss/sitecore-jss-react"] = "22.1.3"

	componentList := components.GetComponents()
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
