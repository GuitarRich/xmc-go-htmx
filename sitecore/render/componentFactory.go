package render

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/guitarrich/headless-go-htmx/model"
)

var components = map[string]interface{}{}

func RegisterComponent(componentName string, component func(model.PlaceholderComponent, model.SitecoreContext) templ.Component) {
	components[componentName] = component
}

func GetComponents() map[string]interface{} {
	return components
}

func GetComponent(component model.PlaceholderComponent, context model.SitecoreContext) templ.Component {
	if component.ComponentName == "" {
		// TODO: Add the "Component does not have an implmementaiton" component
		return templ.Raw("Component Name is empty")
	}

	// Special case for PartialDesignDynamicPlaceholder, this is an internal component
	// and needs to be rendered differently
	if component.ComponentName == "PartialDesignDynamicPlaceholder" {
		return RenderPartialDesignDynamicPlaceholder(component.Params.DynamicPlaceholderID, component, context)
	}

	if components[component.ComponentName] == nil {
		return templ.Raw(fmt.Sprintf("Component [%s] not found", component.ComponentName))
	}

	return components[component.ComponentName].(func(model.PlaceholderComponent, model.SitecoreContext) templ.Component)(component, context)
}
