package editing

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/guitarrich/headless-go-htmx/sitecore"
	"github.com/labstack/echo/v4"
)

func IsEditingMedataPreviewData(data EditingMetadataPreviewData) bool {
	return data.EditMode == EditModeMetadata
}

func (h *EditingRequestHandler) Render(c echo.Context) error {

	fmt.Println("editingRenderMiddlewareRender")
	fmt.Printf("Request: %s\n", c.Request().URL.Path)
	sitecore.EditingLog(
		"editing render middleware start:\n%s \n%s \n%s \n%s\n",
		c.Request().Method,
		c.Request().URL.Query(),
		c.Request().Header,
		c.Request().Body,
	)

	startTimeStamp := time.Now()

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

	// Validate the query string params
	var requiredQueryParams = []string{
		"sc_site",
		"sc_itemid",
		"sc_lang",
		"route",
		"mode",
	}
	var missingQueryParams []string

	for _, requiredQueryParam := range requiredQueryParams {
		if !c.Request().URL.Query().Has(requiredQueryParam) {
			missingQueryParams = append(missingQueryParams, requiredQueryParam)
		}
	}

	if len(missingQueryParams) > 0 {
		message := fmt.Sprintf("Missing query params: %s", missingQueryParams)
		return c.HTML(http.StatusBadRequest, message)
	}

	var previewData EditingMetadataPreviewData
	previewData.Site = c.Request().URL.Query().Get("sc_site")
	previewData.ItemId = c.Request().URL.Query().Get("sc_itemid")
	previewData.Language = c.Request().URL.Query().Get("sc_lang")
	previewData.EditMode = EditModeMetadata
	previewData.PageState = c.Request().URL.Query().Get("mode")
	previewData.VariantIds = strings.Split(c.Request().URL.Query().Get("sc_variant"), ",")
	if len(previewData.VariantIds) == 0 {
		previewData.VariantIds = append(previewData.VariantIds, DEFAULT_VARIANT)
	}
	previewData.Version = c.Request().URL.Query().Get("sc_version")
	previewData.LayoutKind = c.Request().URL.Query().Get("sc_layoutkind")

	routePath := c.Request().URL.Query().Get("route")

	sitecore.EditingLog(
		"ediitng render middleware end in %dms: redirect %s",
		startTimeStamp.Sub(time.Now()).Milliseconds(),
		routePath)

	return c.Redirect(http.StatusFound, routePath)
}
