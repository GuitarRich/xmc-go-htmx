package components

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/guitarrich/headless-go-htmx/model"
)

func GetComponent(component model.PlaceholderComponent) templ.Component {
	funcMap := map[string]templ.Component{
		"Promo":                           Promo(component),
		"PartialDesignDynamicPlaceholder": RenderPartialDesignDynamicPlaceholder(component.Params.DynamicPlaceholderID, component),
		"RichText":                        RichText(component),
		"Image":                           Image(component),
		"Container":                       Container(component),
		"Navigation":                      Navigation(component),
	}

	c := funcMap[component.ComponentName]
	if c == nil {
		fmt.Println("Component not found: " + component.ComponentName)
		return templ.Raw("")
	}
	return c
}
