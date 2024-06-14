package components

import (
	"fmt"
	"reflect"

	"github.com/a-h/templ"
	"github.com/guitarrich/headless-go-htmx/model"
)

func GetComponent(component model.PlaceholderComponent) templ.Component {
	fmt.Printf("GetComponent [%s]\n", component.ComponentName)
	fmt.Printf("Fields is type [%T]\n", component.Fields)

	if component.Fields != nil && reflect.TypeOf(component.Fields).Kind() == reflect.Map {
		fmt.Printf("Fields is a map\n")
	}

	switch component.ComponentName {
	case "Promo":
		return Promo(component)
	case "PartialDesignDynamicPlaceholder":
		return RenderPartialDesignDynamicPlaceholder(component.Params.DynamicPlaceholderID, component)
	case "RichText":
		return RichText(component)
	case "Image":
		return Image(component)
	case "Container":
		return Container(component)
	case "Navigation":
		return Navigation(component)
	default:
		return templ.Raw("Component Not Found")
	}
}
