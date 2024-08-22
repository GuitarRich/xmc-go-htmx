package request

import (
	"encoding/json"
	"fmt"

	"github.com/guitarrich/headless-go-htmx/model"
	"github.com/guitarrich/headless-go-htmx/sitecore"
	"github.com/labstack/echo/v4"
)

func (h *RequestPipelineHandler) HandleNotFound(c echo.Context) error {

	fmt.Println("HandleNotFound")

	siteName := sitecore.GetEnvVar("SITECORE_SITE_NAME")
	language := sitecore.GetEnvVar("SITECORE_LANGUAGE")

	query := sitecore.GetNotFoundPageQuery(siteName, language)
	result := sitecore.RunQuery(query)

	jsonString, _ := json.Marshal(result)

	response := model.NotFoundPageResponse{}
	json.Unmarshal(jsonString, &response)

	/*
		if response.Data.Site.SiteInfo.ErrorHandling.NotFoundPage.Rendered == nil {
			// They have not set a not found page
			return c.NoContent(http.StatusNotFound)
		}
	*/

	h.renderedLayout = response.Data.Site.SiteInfo.ErrorHandling.NotFoundPage.Rendered

	return nil
}
