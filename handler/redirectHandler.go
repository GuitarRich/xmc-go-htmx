package handler

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/guitarrich/headless-go-htmx/model"
	"github.com/guitarrich/headless-go-htmx/sitecore"
	"github.com/labstack/echo/v4"
)

func (h *RequestPipelineHandler) HandleRedirects(c echo.Context) error {

	fmt.Println("HandleRedirects")

	query := sitecore.GetRedirectsForSiteQuery(h.siteName)
	result := sitecore.RunQuery(query)

	jsonString, _ := json.Marshal(result)

	response := model.RedirectResponse{}
	json.Unmarshal(jsonString, &response)

	redirects := response.Data.Site.SiteInfo.Redirects

	for _, redirect := range redirects {
		source := strings.TrimRight(redirect.Pattern, "/")
		match, _ := regexp.MatchString(source, c.Request().URL.Path)
		fmt.Println(source)
		fmt.Println(c.Request().URL.Path)
		fmt.Println(match)

		if match {
			fmt.Println("Redirecting to " + redirect.Target)
			redirectType := getRedirectType(redirect.RedirectType)
			targetUrl := redirect.Target
			if redirect.IsQueryStringPreserved && c.Request().URL.RawQuery != "" {
				targetUrl = targetUrl + "?" + c.Request().URL.RawQuery
			}
			c.Redirect(redirectType, redirect.Target)
			return nil
		}
	}

	return nil
}

func getRedirectType(redirectType string) int {
	switch redirectType {
	case "REDIRECT_301":
		return 301
	case "REDIRECT_302":
		return 302
	}

	return 302
}
