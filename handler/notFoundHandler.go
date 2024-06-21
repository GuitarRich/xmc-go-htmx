package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/guitarrich/headless-go-htmx/model"
	"github.com/guitarrich/headless-go-htmx/sitecore"
	"github.com/guitarrich/headless-go-htmx/view/layout"
	"github.com/labstack/echo/v4"
)

func HandleNotFound(c echo.Context) error {

	fmt.Println("HandleNotFound")

	siteName := sitecore.GetEnvVar("SITECORE_SITE_NAME")
	language := sitecore.GetEnvVar("SITECORE_LANGUAGE")

	query := sitecore.GetNotFoundPageQuery(siteName, language)
	result := sitecore.RunQuery(query)

	jsonString, _ := json.Marshal(result)

	response := model.NotFoundPageResponse{}
	json.Unmarshal(jsonString, &response)

	fmt.Println("response: ", response)
	/*
		if response.Data.Site.SiteInfo.ErrorHandling.NotFoundPage.Rendered == nil {
			// They have not set a not found page
			return c.NoContent(http.StatusNotFound)
		}
	*/

	fmt.Println("Updating dynamic placeholders...")
	var tmp model.PlaceholderComponent
	HandleDynamicPlaceholders(tmp, response.Data.Site.SiteInfo.ErrorHandling.NotFoundPage.Rendered.Sitecore.Route.Placeholders, 1)

	return render(c, http.StatusNotFound, layout.MainLayout(response.Data.Site.SiteInfo.ErrorHandling.NotFoundPage.Rendered.Sitecore.Route))
}
