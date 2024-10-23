// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"fmt"

	"github.com/guitarrich/headless-go-htmx/model"
	"github.com/guitarrich/headless-go-htmx/sitecore/render"
)

type NavigationModel struct {
	Id              string
	Styles          []string
	Href            string
	Querystring     string
	NavigationTitle string
	Children        []NavigationModel
}

func Navigation(props model.PlaceholderComponent) templ.Component {
	fields := props.Fields.([]interface{})[0].(map[string]interface{})

	model := buildNavigationModel(fields)
	return defaultNavigation(props, model)
}

func buildNavigationModel(fields map[string]interface{}) NavigationModel {

	// Build the model from the props
	var model NavigationModel
	model.Id = getSafeString(fields["Id"])
	model.Href = getSafeString(fields["Href"])
	model.Styles = getStyles(fields["Styles"])
	model.Querystring = getSafeString(fields["Querystring"])
	model.NavigationTitle = getNavigationTitle(fields["NavigationTitle"])
	model.Children = []NavigationModel{}

	if fields["Children"] != nil {
		children := fields["Children"].([]interface{})
		for _, child := range children {
			model.Children = append(model.Children, buildNavigationModel(child.(map[string]interface{})))
		}
	}

	return model
}

func getSafeString(field interface{}) string {
	if field == nil {
		fmt.Println("GetSafeString: nil")
		return ""
	}
	return field.(string)
}

func getStyles(field interface{}) []string {
	if field == nil {
		return []string{}
	}
	tmp := field.([]interface{})
	var result = []string{}
	for _, v := range tmp {
		result = append(result, v.(string))
	}

	return result
}

func getNavigationTitle(field interface{}) string {
	if field == nil {
		return ""
	}

	return field.(map[string]interface{})["value"].(string)
}

func defaultNavigation(props model.PlaceholderComponent, model NavigationModel) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var templ_7745c5c3_Var2 = []any{render.DecorateComponent("component navigation", props)}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var2).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/components/navigation.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><div class=\"bg-scwhite md:bg-gradient-sc bg-xl text-scwhite\"><nav class=\"w-full shadow-md\"><div><div class=\"hidden md:block\"><div class=\"mx-auto flex max-w-[1004px] justify-between p-4\"></div><div class=\"bg-scwhite text-scblack\"><div class=\"item-center relative mx-auto box-border flex w-[1004px] max-w-full flex-wrap justify-start px-5\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, item := range model.Children {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a href=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 templ.SafeURL = templ.URL(item.Href)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var4)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"flex flex-auto\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(item.NavigationTitle)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/components/navigation.templ`, Line: 89, Col: 32}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div></div></nav></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
