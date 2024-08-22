package request

import (
	"encoding/json"
	"fmt"

	"github.com/guitarrich/headless-go-htmx/model"
	"github.com/guitarrich/headless-go-htmx/sitecore"
	"github.com/labstack/echo/v4"
)

func (h *RequestPipelineHandler) HandleLayout(c echo.Context) error {

	fmt.Println("MainLayoutHandler")
	fmt.Printf("Request: %s\n", c.Request().URL.Path)

	query := sitecore.GetLayoutQuery(h.itemPath, h.language, h.siteName)
	result := sitecore.RunQuery(query)

	jsonString, _ := json.Marshal(result)

	layoutResponse := model.LayoutResponse{}
	json.Unmarshal(jsonString, &layoutResponse)

	if layoutResponse.Data.Layout.Item.Rendered.Sitecore.Route.Placeholders == nil {
		// This is a 404, so we need to render the 404 page
		fmt.Println("NotFound")
		h.NotFound = true
		return nil
	}

	h.renderedLayout = layoutResponse.Data.Layout.Item.Rendered

	return nil
}
