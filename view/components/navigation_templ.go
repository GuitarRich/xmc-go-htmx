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
)

type NavigationModel struct {
	Id             string
	Styles         []string
	Href           string
	Querystring    string
	NavitaionTitle string
	Children       []NavigationModel
}

func Navigation(props model.PlaceholderComponent) templ.Component {
	fmt.Println("Navigation[" + props.ComponentName + "]")
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
	model.NavitaionTitle = getNavigationTitle(fields["NavigationTitle"])
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<nav class=\"w-full flex flex-wrap relative bg-gray-800 text-white\"><div class=\"w-full flex flex-wrap relative\"><ul class=\"w-full flex flex-wrap relative\"><li class=\"w-full flex flex-wrap relative\"><a href=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 templ.SafeURL = templ.SafeURL(model.Href)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var2)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"w-full flex flex-wrap relative\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(model.NavitaionTitle)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/components/navigation.templ`, Line: 80, Col: 104}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></li>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, child := range model.Children {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"w-full flex flex-wrap relative\"><a href=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 templ.SafeURL = templ.SafeURL(child.Href)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var4)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"w-full flex flex-wrap relative\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(child.NavitaionTitle)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/components/navigation.templ`, Line: 84, Col: 105}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul></div></nav>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
