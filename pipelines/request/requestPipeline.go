package request

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/guitarrich/headless-go-htmx/middleware/editing"
	"github.com/guitarrich/headless-go-htmx/model"
	"github.com/guitarrich/headless-go-htmx/sitecore"
	"github.com/labstack/echo/v4"
)

func (h *RequestPipelineHandler) Run(c echo.Context) error {

	sitecore.RendererLog("RequestPipelineHandler\nRequest: %s\n", c.Request().URL.Path)

	// are we in preview mode?
	var query string
	var queryName string
	headers := map[string]string{}
	params := map[string]string{}

	sitecore.RendererLog("Query: %s\n", c.Request().URL.RawQuery)
	sitecore.RendererLog("mode: %s\n", c.Request().URL.Query().Get("mode"))

	editMode := c.Request().URL.Query().Get("mode") == "edit"
	if editMode {
		sitecore.RendererLog("This is Editing mode\n")
		queryName = "EditingQuery"
		var previewData editing.EditingMetadataPreviewData
		previewData.Site = c.Request().URL.Query().Get("sc_site")
		previewData.ItemId = c.Request().URL.Query().Get("sc_itemid")
		previewData.Language = c.Request().URL.Query().Get("sc_lang")
		previewData.EditMode = editing.EditModeMetadata
		previewData.PageState = c.Request().URL.Query().Get("mode")
		previewData.VariantIds = strings.Split(c.Request().URL.Query().Get("sc_variant"), ",")
		previewData.Version = c.Request().URL.Query().Get("sc_version")
		previewData.LayoutKind = c.Request().URL.Query().Get("sc_layoutKind")

		fmt.Printf("layoutKind: %s\n", previewData.LayoutKind)

		query = sitecore.GetEditingDataQuery()

		headers["sc_layoutKind"] = previewData.LayoutKind
		headers["sc_editmode"] = "true"

		params["siteName"] = previewData.Site
		params["itemId"] = previewData.ItemId
		params["language"] = previewData.Language
		params["version"] = previewData.Version
		params["after"] = ""
		params["pageSize"] = "50"

		result := sitecore.RunQueryWithParameters(query, queryName, headers, params)

		jsonString, _ := json.Marshal(result)

		layoutResponse := model.EditingResponse{}
		json.Unmarshal(jsonString, &layoutResponse)

		h.renderedLayout = layoutResponse.Data.Item.Rendered
	} else {
		sitecore.RendererLog("Normal mode\n")
		queryName = "LayoutQuery"
		query = sitecore.GetLayoutQuery(h.itemPath, h.language, h.siteName)

		result := sitecore.RunQueryWithParameters(query, queryName, headers, params)

		jsonString, _ := json.Marshal(result)

		layoutResponse := model.LayoutResponse{}
		json.Unmarshal(jsonString, &layoutResponse)

		if layoutResponse.Data.Layout.Item.Rendered.Sitecore.Route.Placeholders == nil {
			// This is a 404, so we need to render the 404 page
			sitecore.RendererLog("404: NotFound\n")
			h.NotFound = true
			return nil
		}

		h.renderedLayout = layoutResponse.Data.Layout.Item.Rendered
	}

	return nil
}
