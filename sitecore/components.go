package sitecore

import (
	"context"
	"fmt"
	"io"

	//	"reflect"

	"github.com/a-h/templ"
	"github.com/guitarrich/headless-go-htmx/model"
)

func DecorateComponent(cssClass string, props model.PlaceholderComponent) string {
	return fmt.Sprintf("%s %s %s", cssClass, props.Params.GridParameters, props.Params.Styles)
}

func RenderRichText(field model.RichTextField) templ.Component {
	fmt.Println("   --> RenderRichText")
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, fmt.Sprintf("%s", field.Value))
		return err
	})
}

func GetRichTextField(fields interface{}, fieldName string) model.RichTextField {
	fieldMap, ok := fields.(map[string]interface{})
	if !ok {
		fmt.Println("GetRichTextField: not a map")
		return model.RichTextField{}
	}
	baseField, ok := fieldMap[fieldName].(map[string]interface{})["value"]
	if !ok {
		fmt.Println("GetRichTextField: not a field")
		return model.RichTextField{}
	}

	var richTextField model.RichTextField
	richTextField.Value = getSafeString(baseField)
	return richTextField
}

func RenderLink(field model.LinkField) templ.Component {
	fmt.Println("   --> RenderLink")

	href := field.Href
	if field.Querystring != "" {
		href += "?" + field.Querystring
	}
	if field.Anchor != "" {
		href += "#" + field.Anchor
	}

	link := fmt.Sprintf("<a  href=\"%s\"", href)
	if field.Target != "" {
		link += fmt.Sprintf(" target=\"%s\"", field.Target)
	}
	if field.Title != "" {
		link += fmt.Sprintf(" title=\"%s\"", field.Title)
	}
	if field.Class != "" {
		link += fmt.Sprintf(" class=\"%s\"", field.Class)
	}

	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, link)
		return err
	})
}

func GetLinkField(fields interface{}, fieldName string) model.LinkField {
	fieldMap, ok := fields.(map[string]interface{})
	if !ok {
		fmt.Println("GetLinkField: not a map")
		return model.LinkField{}
	}
	baseField, ok := fieldMap[fieldName].(map[string]interface{})["value"].(map[string]interface{})
	if !ok {
		fmt.Println("GetLinkField: not a field")
		return model.LinkField{}
	}

	var linkField model.LinkField
	linkField.Text = getSafeString(baseField["text"])
	linkField.Href = getSafeString(baseField["href"])
	linkField.Anchor = getSafeString(baseField["anchor"])
	linkField.Querystring = getSafeString(baseField["querystring"])
	linkField.Target = getSafeString(baseField["target"])
	linkField.Title = getSafeString(baseField["title"])
	linkField.Class = getSafeString(baseField["class"])

	return linkField
}

func RenderImage(field model.ImageField) templ.Component {
	fmt.Println("   --> RenderImage")

	img := fmt.Sprintf("<img src=\"%s\"", field.Src)
	if field.Alt != "" {
		img += fmt.Sprintf(" alt=\"%s\"", field.Alt)
	}
	if field.Width != "" {
		img += fmt.Sprintf(" width=\"%s\"", field.Width)
	}
	if field.Height != "" {
		img += fmt.Sprintf(" height=\"%s\"", field.Height)
	}

	img += fmt.Sprintf(" />")

	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, img)
		return err
	})
}

func getSafeString(field interface{}) string {
	if field == nil {
		return ""
	}
	return fmt.Sprintf("%s", field)
}

func GetImageField(fields interface{}, fieldName string) model.ImageField {
	fieldMap, ok := fields.(map[string]interface{})
	if !ok {
		fmt.Println("GetImageField: not a map")
		return model.ImageField{}
	}
	baseField, ok := fieldMap[fieldName].(map[string]interface{})["value"].(map[string]interface{})
	if !ok {
		fmt.Println("GetImageField: not a field")
		return model.ImageField{}
	}

	var imageField model.ImageField
	imageField.Src = getSafeString(baseField["src"])
	imageField.Alt = getSafeString(baseField["alt"])
	imageField.Width = getSafeString(baseField["width"])
	imageField.Height = getSafeString(baseField["height"])

	return imageField
}
