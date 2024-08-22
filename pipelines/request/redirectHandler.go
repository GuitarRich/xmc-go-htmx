package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/allegro/bigcache"
	"github.com/guitarrich/headless-go-htmx/model"
	"github.com/guitarrich/headless-go-htmx/sitecore"
	"github.com/labstack/echo/v4"
)

func (h *RequestPipelineHandler) HandleRedirects(c echo.Context) error {

	fmt.Println("HandleRedirects")

	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(100 * time.Minute))

	fmt.Println(" -> Checking cache")
	jsonString, err := cache.Get(fmt.Sprintf("redirects-%s", h.siteName))
	if err != nil {
		fmt.Println(" -> Cache miss")
		query := sitecore.GetRedirectsForSiteQuery(h.siteName)
		result := sitecore.RunQuery(query)

		jsonString, _ = json.Marshal(result)
		cache.Set(fmt.Sprintf("redirects-%s", h.siteName), jsonString)
	} else {
		fmt.Println(" -> Cache hit")
	}

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
		return http.StatusMovedPermanently
	case "REDIRECT_302":
		return http.StatusFound
	}

	return http.StatusFound
}
