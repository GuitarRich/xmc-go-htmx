package render

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/a-h/templ"
	"github.com/guitarrich/headless-go-htmx/model"
	"github.com/guitarrich/headless-go-htmx/sitecore"
	"github.com/mitchellh/mapstructure"
)

func DecorateComponent(cssClass string, props model.PlaceholderComponent) string {
	return fmt.Sprintf("%s %s %s", cssClass, props.Params.GridParameters, props.Params.Styles)
}

func renderFieldMetadata(metadata model.MetadataData, field templ.Component) templ.Component {
	fmt.Printf("renderFieldMetadata: %T, %#v\n", metadata, metadata)
	fieldMetadata := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return RenderFieldMetadata(metadata, field).Render(ctx, w)
	})
	return fieldMetadata
}

func RichTextField(fields interface{}, fieldName string) templ.Component {
	return renderRichText(GetRichTextField(fields, fieldName))
}

func renderRichText(field model.RichTextField) templ.Component {
	renderedField := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, fmt.Sprintf("%s", field.Value))
		return err
	})

	return RenderFieldMetadata(field.Metadata, renderedField)
}

func GetRichTextField(fields interface{}, fieldName string) model.RichTextField {
	fieldMap, ok := fields.(map[string]interface{})
	if !ok {
		fmt.Println("GetRichTextField: not a map")
		return model.RichTextField{}
	}
	baseField, ok := fieldMap[fieldName].(map[string]interface{})
	if !ok {
		fmt.Println("GetRichTextField: not a field")
		return model.RichTextField{}
	}

	var result model.RichTextField
	err := mapstructure.Decode(baseField, &result)
	if err != nil {
		fmt.Printf("GetRichTextField: not a RichTextField, %s", err)
		return model.RichTextField{}
	}

	return result
}

func LinkField(fields interface{}, fieldName string) templ.Component {
	field := GetLinkField(fields, fieldName)
	return renderFieldMetadata(field.Metadata, renderLink(field))
}

func renderLink(field model.LinkField) templ.Component {
	href := field.Href
	if field.Querystring != "" {
		href += "?" + field.Querystring
	}
	if field.Anchor != "" {
		href += "#" + field.Anchor
	}

	link := fmt.Sprintf("<a  href=\"%s\"", href)
	link += sitecore.AddIfNotEmpty("target", field.Target)
	link += sitecore.AddIfNotEmpty("title", field.Title)
	link += sitecore.AddIfNotEmpty("class", field.Class)

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

	baseField, ok := fieldMap[fieldName].(map[string]interface{})
	if !ok {
		fmt.Println("GetLinkField: not a field")
		return model.LinkField{}
	}

	var result model.LinkField

	err := mapstructure.Decode(baseField, &result)
	if err != nil {
		fmt.Printf("GetLinkField: not a LinkField, %s", err)
		return model.LinkField{}
	}

	return result
}

func ImageField(fields interface{}, fieldName string) templ.Component {
	field := GetImageField(fields, fieldName)
	return renderFieldMetadata(field.Metadata, renderImage(field))
}

func renderImage(field model.ImageField) templ.Component {
	img := fmt.Sprintf("<img src=\"%s\"", field.Value.Src)
	img += sitecore.AddIfNotEmpty("alt", field.Value.Alt)
	img += sitecore.AddIfNotEmpty("width", field.Value.Width)
	img += sitecore.AddIfNotEmpty("height", field.Value.Height)

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
	imageField.Value.Src = sitecore.GetSafeString(baseField["src"])
	imageField.Value.Alt = sitecore.GetSafeString(baseField["alt"])
	imageField.Value.Width = sitecore.GetSafeString(baseField["width"])
	imageField.Value.Height = sitecore.GetSafeString(baseField["height"])

	return imageField
}

func RenderPlaceholderOpen(placeholderKey string, id string, context model.SitecoreContext) string {
	isEditing := sitecore.IsEditMode(context)
	if id == "" {
		id = "00000000-0000-0000-0000-000000000000"
	}
	var result strings.Builder
	if isEditing {
		result.WriteString("<code type=\"text/sitecore\" chrometype=\"placeholder\" class=\"scpm\" kind=\"open\"")
		result.WriteString("id=\"")
		result.WriteString(fmt.Sprintf("%s-%s", placeholderKey, id))
		result.WriteString("\" style=\"cursor:pointer;\"></code>")
	}

	return result.String()
}

func RenderPlaceholderClose(placeholderKey string, context model.SitecoreContext) string {
	isEditing := sitecore.IsEditMode(context)
	var result strings.Builder
	if isEditing {
		result.WriteString("<code type=\"text/sitecore\" chrometype=\"placeholder\" class=\"scpm\" kind=\"close\" style=\"cursor:pointer;\"></code>")
	}

	return result.String()
}

func RenderComponentOpen(component model.PlaceholderComponent, context model.SitecoreContext) string {
	isEditing := sitecore.IsEditMode(context)
	var result strings.Builder
	if isEditing {
		result.WriteString("<code type=\"text/sitecore\" chrometype=\"rendering\" class=\"scpm\" kind=\"open\"")
		result.WriteString("id=\"")
		result.WriteString(component.UID)
		result.WriteString("\" style=\"cursor:pointer;\"></code>")
	}

	return result.String()
}

func RenderComponentClose(component model.PlaceholderComponent, context model.SitecoreContext) string {
	isEditing := sitecore.IsEditMode(context)
	var result strings.Builder
	if isEditing {
		result.WriteString("<code type=\"text/sitecore\" chrometype=\"rendering\" class=\"scpm\" kind=\"close\" style=\"cursor:pointer;\"></code>")
	}

	return result.String()
}
