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
	fieldMetadata := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return RenderFieldMetadata(metadata, field).Render(ctx, w)
	})
	return fieldMetadata
}

func RichTextField(fields interface{}, fieldName string) templ.Component {
	field, _ := GetRichTextField(fields, fieldName)
	return RenderRichText(field)
}

func RenderRichText(field model.RichTextField) templ.Component {
	renderedField := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, fmt.Sprintf("%s", field.Value))
		return err
	})

	return RenderFieldMetadata(field.Metadata, renderedField)
}

func GetFieldWithFallback(fields interface{}, fieldNames ...string) string {
	var result model.RichTextField
	var ok bool
	for _, fieldName := range fieldNames {
		result, ok = GetRichTextField(fields, fieldName)
		if ok {
			return fmt.Sprintf("%s", result.Value)
		}
	}

	return ""
}

func GetRichTextField(fields interface{}, fieldName string) (model.RichTextField, bool) {
	fieldMap, ok := fields.(map[string]interface{})
	if !ok {
		fmt.Println("GetRichTextField: not a map")
		return model.RichTextField{}, false
	}
	baseField, ok := fieldMap[fieldName].(map[string]interface{})
	if !ok {
		fmt.Println("GetRichTextField: not a field")
		return model.RichTextField{}, false
	}

	var result model.RichTextField
	err := mapstructure.Decode(baseField, &result)
	if err != nil {
		fmt.Printf("GetRichTextField: not a RichTextField, %s", err)
		return model.RichTextField{}, false
	}

	return result, true
}

func LinkField(fields interface{}, fieldName string, classNames ...string) templ.Component {
	field := GetLinkField(fields, fieldName)
	return renderFieldMetadata(field.Metadata, renderLink(field, classNames...))
}

func LinkFieldHasLink(fields interface{}, fieldName string) bool {
	field := GetLinkField(fields, fieldName)
	return field.Value.Href != ""
}

func renderLink(field model.LinkField, classNames ...string) templ.Component {
	href := field.Value.Href
	if field.Value.Querystring != "" {
		href += "?" + field.Value.Querystring
	}
	if field.Value.Anchor != "" {
		href += "#" + field.Value.Anchor
	}

	link := fmt.Sprintf("<a  href=\"%s\"", href)
	link += sitecore.AddIfNotEmpty("target", field.Value.Target)
	link += sitecore.AddIfNotEmpty("title", field.Value.Title)

	cssClasses := field.Value.Class
	for _, className := range classNames {
		cssClasses += " " + className
	}
	link += sitecore.AddIfNotEmpty("class", cssClasses)

	if field.Value.Target == "_blank" {
		link += " rel=\"noopener noreferrer\""
	}

	link += fmt.Sprintf(">%s</a>", field.Value.Text)

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

func ImageField(fields interface{}, fieldName string, cssClass string) templ.Component {
	field := GetImageField(fields, fieldName)
	return renderFieldMetadata(field.Metadata, renderImage(field, cssClass))
}

func renderImage(field model.ImageField, cssClass string) templ.Component {
	img := fmt.Sprintf("<img src=\"%s\"", field.Value.Src)
	img += sitecore.AddIfNotEmpty("alt", field.Value.Alt)
	img += sitecore.AddIfNotEmpty("width", field.Value.Width)
	img += sitecore.AddIfNotEmpty("height", field.Value.Height)

	if cssClass != "" {
		img += fmt.Sprintf(" class=\"%s\"", cssClass)
	}

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
	baseField, ok := fieldMap[fieldName].(map[string]interface{})
	if !ok {
		fmt.Println("GetImageField: not a field")
		return model.ImageField{}
	}

	var imageField model.ImageField
	err := mapstructure.Decode(baseField, &imageField)
	if err != nil {
		fmt.Printf("GetImageField: not an ImageField, %s", err)
		return model.ImageField{}
	}

	return imageField
}

func RenderPlaceholderOpen(placeholderKey string, id string, context model.SitecoreContext) string {
	if id == "" {
		id = "00000000-0000-0000-0000-000000000000"
	}
	var result strings.Builder
	if context.PageEditing {
		result.WriteString("<code type=\"text/sitecore\" chrometype=\"placeholder\" class=\"scpm\" kind=\"open\"")
		result.WriteString("id=\"")
		result.WriteString(fmt.Sprintf("%s_%s", placeholderKey, id))
		result.WriteString("\" style=\"cursor:pointer;\"></code>")
	}

	return result.String()
}

func RenderPlaceholderClose(placeholderKey string, context model.SitecoreContext) string {
	if context.PageEditing {
		return "<code type=\"text/sitecore\" chrometype=\"placeholder\" class=\"scpm\" kind=\"close\" style=\"cursor:pointer;\"></code>"
	}

	return ""
}

func RenderComponentOpen(component model.PlaceholderComponent, context model.SitecoreContext) string {
	var result strings.Builder
	if context.PageEditing {
		result.WriteString("<code type=\"text/sitecore\" chrometype=\"rendering\" class=\"scpm\" kind=\"open\"")
		result.WriteString("id=\"")
		result.WriteString(component.UID)
		result.WriteString("\" style=\"cursor:pointer;\"></code>")
	}

	return result.String()
}

func RenderComponentClose(component model.PlaceholderComponent, context model.SitecoreContext) string {
	if context.PageEditing {
		return "<code type=\"text/sitecore\" chrometype=\"rendering\" class=\"scpm\" kind=\"close\" style=\"cursor:pointer;\"></code>"
	}

	return ""
}
