package request

import (
	"net/http"

	"github.com/guitarrich/headless-go-htmx/model"
	"github.com/guitarrich/headless-go-htmx/sitecore"
	"github.com/guitarrich/headless-go-htmx/view/layout"
	"github.com/labstack/echo/v4"
)

type RequestPipelineHandler struct {
	siteName       string
	language       string
	itemPath       string
	itemId         string
	context        echo.Context
	renderedLayout model.Rendered
	NotFound       bool
}

func (h RequestPipelineHandler) RequestBeginHandler(c echo.Context) error {

	h.siteName = sitecore.GetEnvVar("SITECORE_SITE_NAME")
	h.language = sitecore.GetEnvVar("SITECORE_LANGUAGE")
	h.itemPath = c.Request().URL.Path

	// Handle special cases
	if h.itemPath == "/sitemap.xml" {
		return h.HandleSitemap(c)
	}

	// Handle cathall paths
	_ = h.HandleLayout(c)
	if h.NotFound {
		// Check for redirects
		_ = h.HandleRedirects(c)
		if h.NotFound {
			_ = h.HandleNotFound(c)
		}
	}

	var tmp model.PlaceholderComponent
	HandleDynamicPlaceholders(tmp, h.renderedLayout.Sitecore.Route.Placeholders, 1)

	response := render(c, http.StatusOK, layout.MainLayout(h.renderedLayout.Sitecore.Route))
	return response
}
