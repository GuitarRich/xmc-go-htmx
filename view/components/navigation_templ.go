// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"encoding/xml"
	"fmt"

	"github.com/guitarrich/headless-go-htmx/model"
	"github.com/guitarrich/headless-go-htmx/sitecore/render"
	"github.com/mitchellh/mapstructure"
)

type NavigationModel struct {
	navigationLinks []NavigationLink
}

type NavigationTitle struct {
	Value    string             `json:"value"`
	Editable string             `json:"editable"`
	Metadata model.MetadataData `json:"metadata"`
}

type NavigationLink struct {
	Id              string
	Styles          []string
	Href            string
	Querystring     string
	NavigationTitle model.RichTextField
	Children        []NavigationLink
}

func Navigation(props model.PlaceholderComponent, sc model.SitecoreContext) templ.Component {
	navigationLinks := props.Fields.([]interface{})

	var model NavigationModel
	for _, navigationLink := range navigationLinks {
		model.navigationLinks = append(model.navigationLinks, buildNavigationModel(navigationLink.(map[string]interface{})))
	}
	return defaultNavigation(props, model)
}

func buildNavigationModel(fields map[string]interface{}) NavigationLink {
	var model NavigationLink
	err := mapstructure.Decode(fields, &model)
	if err != nil {
		fmt.Printf("buildNavigationModel: not a NavigationLink, %s", err)
		return model
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

func getBackgroundImage(props model.PlaceholderComponent) string {
	if props.Fields == nil {
		return ""
	}
	var imageFragment model.ImageFragment
	err := xml.Unmarshal([]byte(props.Params.BackgroundImage), &imageFragment)
	if err != nil {
		fmt.Printf("getBackgroundImage: not an imageFragment, %s", err)
		return ""
	}
	return imageFragment.Src
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><div class=\"bg-scwhite bg-xl text-scwhite md:bg-gradient-sc\"><nav class=\"w-full shadow-md\"><div><div class=\"hidden md:block\"><div class=\"mx-auto flex max-w-[1004px] justify-between p-4\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, item := range model.navigationLinks {
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
			templ_7745c5c3_Err = render.RenderRichText(item.NavigationTitle).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"bg-scwhite text-scblack\"><div class=\"item-center relative mx-auto box-border flex w-[1004px] max-w-full flex-wrap justify-start px-5\"><div class=\"flex basis-1/6 flex-wrap items-center justify-start pb-4 pt-2\"><a href=\"/\" class=\"flex-shrink-0\"><img src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(getBackgroundImage(props))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/components/navigation.templ`, Line: 113, Col: 46}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" alt=\"Logo\" class=\"auto h-8\"></a></div><div class=\"flex basis-5/6 flex-wrap items-stretch justify-end space-x-4 pl-12\"><a href=\"/search\" class=\"justify-end text-scblack my-4\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" class=\"lucide lucide-search h-5 w-5\"><circle cx=\"11\" cy=\"11\" r=\"8\"></circle> <path d=\"m21 21-4.3-4.3\"></path></svg></a></div></div></div></div><div class=\"block md:hidden\"><div class=\"flex h-12 items-center justify-between px-4 pb-5 pt-4 md:px-0\"><a href=\"/\" class=\"mx-auto flex-shrink-0\"><img src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(getBackgroundImage(props))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/components/navigation.templ`, Line: 141, Col: 44}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" alt=\"Logo\" class=\"h-8 w-auto\"></a> <button class=\"z-50 text-scblack hover:text-scblue\"><svg class=\"h-6 w-6\" xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" class=\"lucide lucide-menu h-6 w-6\"><line x1=\"4\" x2=\"20\" y1=\"12\" y2=\"12\"></line> <line x1=\"4\" x2=\"20\" y1=\"6\" y2=\"6\"></line> <line x1=\"4\" x2=\"20\" y1=\"18\" y2=\"18\"></line></svg></button></div></div></div></nav></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
