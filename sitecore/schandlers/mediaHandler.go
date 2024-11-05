package schandlers

import (
	"fmt"
	"net/http"

	"github.com/guitarrich/headless-go-htmx/sitecore"
	"github.com/labstack/echo/v4"
)

type MediaHandler struct{}

func NewMediaHandler() *MediaHandler {
	return &MediaHandler{}
}

func (h *MediaHandler) ServeHTTP(c echo.Context) error {
	// This handler should only happen in edit mode or when pointing at the
	// preview end point. So we just need to redirect the request through to the
	// CM. Let's use the CONTEXT_ID to get the CM Url

	sitecoreApiHost := sitecore.GetEnvVar("SITECORE_API_HOST")

	c.Redirect(http.StatusFound, fmt.Sprintf("%s/%s", sitecoreApiHost, c.Request().URL.Path))

	return nil
}
