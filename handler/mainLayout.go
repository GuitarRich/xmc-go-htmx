package handler

import (
	"encoding/json"
	"fmt"
	"strings"

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

	fmt.Println("itemPath: " + itemPath)
	if itemPath == "/sitemap.xml" {
		// todo: we need to stream the sitemap.xml file from Edge
		return HandleSitemap(c)
	}

	query := sitecore.GetLayoutQuery(itemPath, language, siteName)
	result := sitecore.RunQuery(query)

	jsonString, _ := json.Marshal(result)

	layoutResponse := model.LayoutResponse{}
	json.Unmarshal(jsonString, &layoutResponse)

	fmt.Println("Updating dynamic placeholders...")
	var tmp model.PlaceholderComponent
	HandleDynamicPlaceholders(tmp, layoutResponse.Data.Layout.Item.Rendered.Sitecore.Route.Placeholders, 1)

	return render(c, layout.MainLayout(layoutResponse.Data.Layout.Item.Rendered.Sitecore.Route))
}

func HandleDynamicPlaceholders(component model.PlaceholderComponent, placeholders map[string][]model.PlaceholderComponent, level int) {

	for key, val := range placeholders {
		fmt.Printf(" ==> PlaceholderKey: [%s]\n", key)
		if strings.HasSuffix(key, "-{*}") {
			fmt.Println("   -> Has dynamic placeholder")
			newKey := strings.Replace(key, "{*}", component.Params.DynamicPlaceholderID, -1)
			placeholders[newKey] = val
		}
		for _, component := range val {
			if len(component.Placeholders) > 0 {
				fmt.Println("   -> Has nested placholders")
				HandleDynamicPlaceholders(component, component.Placeholders, level+1)
			}
		}
	}
}
