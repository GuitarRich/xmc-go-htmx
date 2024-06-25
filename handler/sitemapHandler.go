package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/a-h/templ"
	"github.com/guitarrich/headless-go-htmx/model"
	"github.com/guitarrich/headless-go-htmx/sitecore"
	"github.com/labstack/echo/v4"
)

func (h *RequestPipelineHandler) HandleSitemap(c echo.Context) error {

	siteName := sitecore.GetEnvVar("SITECORE_SITE_NAME")
	itemPath := c.Request().URL.Path

	fmt.Println("itemPath: " + itemPath)
	if itemPath != "/sitemap.xml" {
		// we need to throw an error here as this should only be called for /sitemap.xml
		return c.NoContent(http.StatusNotFound)
	}

	query := sitecore.GetSitemapQuery(siteName)
	result := sitecore.RunQuery(query)

	jsonString, _ := json.Marshal(result)
	fmt.Println("jsonString: " + string(jsonString))

	resposne := model.SitemapResponse{}
	json.Unmarshal(jsonString, &resposne)
	fmt.Println("resposne: ", resposne)

	c.Response().Header().Set("Content-Type", "application/xml")
	c.Response().WriteHeader(http.StatusOK)

	fmt.Println("Fetching sitemap from: ", resposne.Data.Site.SiteInfo.Sitemap)
	for _, url := range resposne.Data.Site.SiteInfo.Sitemap {
		fmt.Println("Fetching: ", url)
		sitemap, _ := getXml(url)

		fmt.Println("sitemap: " + string(sitemap))
		return render(c, http.StatusOK, templ.Raw(string(sitemap)))
	}

	return c.NoContent(http.StatusOK)
}

func getXml(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("error fetching %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("error fetching %s: status code %d", url, resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("error reading response body: %w", err)
	}

	return data, nil
}
