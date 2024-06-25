package handler

import (
	"strings"

	"github.com/a-h/templ"
	"github.com/guitarrich/headless-go-htmx/model"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, statusCode int, component templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := component.Render(c.Request().Context(), buf); err != nil {
		return err
	}

	return c.HTML(statusCode, buf.String())
	//return component.Render(c.Request().Context(), c.Response())
}

func HandleDynamicPlaceholders(component model.PlaceholderComponent, placeholders map[string][]model.PlaceholderComponent, level int) {

	for key, val := range placeholders {
		if strings.HasSuffix(key, "-{*}") {
			newKey := strings.Replace(key, "{*}", component.Params.DynamicPlaceholderID, -1)
			placeholders[newKey] = val
		}
		for _, component := range val {
			if len(component.Placeholders) > 0 {
				HandleDynamicPlaceholders(component, component.Placeholders, level+1)
			}
		}
	}
}
