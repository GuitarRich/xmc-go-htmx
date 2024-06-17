package components

import (
	"fmt"
	"reflect"

	"github.com/a-h/templ"
	"github.com/guitarrich/headless-go-htmx/model"
)

var components = map[string]interface{}{}

func RegisterComponent(componentName string, component func(model.PlaceholderComponent) templ.Component) {
	components[componentName] = component
}

func GetComponent(component model.PlaceholderComponent) templ.Component {
	fmt.Printf("GetComponent [%s]\n", component.ComponentName)
	fmt.Printf("Fields is type [%T]\n", component.Fields)

	if component.Fields != nil && reflect.TypeOf(component.Fields).Kind() == reflect.Map {
		fmt.Printf("Fields is a map\n")
	}

	if component.ComponentName == "" {
		return templ.Raw("Component Name is empty")
	}

	// Special case for PartialDesignDynamicPlaceholder, this is an internal component
	// and needs to be rendered differently
	if component.ComponentName == "PartialDesignDynamicPlaceholder" {
		return RenderPartialDesignDynamicPlaceholder(component.Params.DynamicPlaceholderID, component)
	}

	if components[component.ComponentName] == nil {
		return templ.Raw(fmt.Sprintf("Component [%s] not found", component.ComponentName))
	}

	return components[component.ComponentName].(func(model.PlaceholderComponent) templ.Component)(component)
}
