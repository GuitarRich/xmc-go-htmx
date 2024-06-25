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
	href := field.Href
	if field.Querystring != "" {
		href += "?" + field.Querystring
	}
	if field.Anchor != "" {
		href += "#" + field.Anchor
	}

	link := fmt.Sprintf("<a  href=\"%s\"", href)
	link += AddIfNotEmpty("target", field.Target)
	link += AddIfNotEmpty("title", field.Title)
	link += AddIfNotEmpty("class", field.Class)

	if field.Target == "_blank" {
		link += " rel=\"noopener noreferrer\""
	}

	link += fmt.Sprintf(">%s</a>", field.Text)

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
	fmt.Println(fieldMap[fieldName])
	baseField, ok := fieldMap[fieldName].(map[string]interface{})["value"].(map[string]interface{})
	fmt.Println(baseField)
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
	img := fmt.Sprintf("<img src=\"%s\"", field.Src)
	img += AddIfNotEmpty("alt", field.Alt)
	img += AddIfNotEmpty("width", field.Width)
	img += AddIfNotEmpty("height", field.Height)

	img += fmt.Sprintf(" />")

	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, img)
		return err
	})
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
