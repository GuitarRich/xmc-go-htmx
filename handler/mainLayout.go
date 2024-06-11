package handler

import (
	"encoding/json"
	"fmt"

	"github.com/guitarrich/headless-go-htmx/model"
	"github.com/guitarrich/headless-go-htmx/sitecore"
	"github.com/guitarrich/headless-go-htmx/view/layout"
	"github.com/labstack/echo/v4"
)

type MainLayoutHandler struct{}

func (h MainLayoutHandler) HandleLayout(c echo.Context) error {

	fmt.Println("MainLayoutHandler")
	fmt.Println(c.Request().URL.Path)

	siteName := sitecore.GetEnvVar("SITECORE_SITE_NAME")
	language := sitecore.GetEnvVar("SITECORE_LANGUAGE")
	itemPath := c.Request().URL.Path

	query := sitecore.GetLayoutQuery(itemPath, language, siteName)
	result := sitecore.RunQuery(query)

	jsonString, _ := json.Marshal(result)
	fmt.Println(string(jsonString))

	layoutResponse := model.LayoutResponse{}
	json.Unmarshal(jsonString, &layoutResponse)

	fmt.Println("Placeholders")
	for key, val := range layoutResponse.Data.Layout.Item.Rendered.Sitecore.Route.Placeholders {
		fmt.Println(key, val)
		fmt.Println(" -> Components:")

		for i, component := range val {
			fmt.Println(i, component)
		}
	}

	return render(c, layout.MainLayout(layoutResponse.Data.Layout.Item.Rendered.Sitecore.Route))
}
