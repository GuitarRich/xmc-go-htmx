package sitecore

import (
	"context"
	"fmt"
	"io"
	"reflect"

	"github.com/a-h/templ"
	"github.com/guitarrich/headless-go-htmx/model"
)

func DecorateComponent(cssClass string, props model.PlaceholderComponent) string {
	return fmt.Sprintf("%s %s %s", cssClass, props.Params.GridParameters, props.Params.Styles)
}

func RenderRichText(field interface{}) templ.Component {
	fmt.Println("RenderRichText")
	fmt.Println(field)
	scField, ok := field.(model.ScField)
	if !ok {
		fmt.Println("RenderRichText: not a map")
		return nil
	}
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, fmt.Sprintf("%s", scField.Value))
		return err
	})
}

func RenderLink(field interface{}) templ.Component {
	fmt.Println("RenderLink")
	fmt.Println(field)
	scField, ok := field.(model.ScField)
	if !ok {
		fmt.Println("RenderLink: not a ScField")
		return nil
	}
	linkField, ok := scField.Value.(map[string]interface{})
	if !ok {
		fmt.Println("RenderLink: not a linkField")
		return nil
	}

	anchor := linkField["anchor"].(string)
	querystring := linkField["querystring"].(string)
	href := linkField["href"].(string)
	text := linkField["text"].(string)
	target := linkField["target"].(string)
	title := linkField["title"].(string)
	class := linkField["class"].(string)

	if querystring != "" {
		href += "?" + querystring
	}
	if anchor != "" {
		href += "#" + anchor
	}

	link := fmt.Sprintf("<a  href=\"%s\"", href)
	if target != "" {
		link += fmt.Sprintf(" target=\"%s\"", target)
	}
	if title != "" {
		link += fmt.Sprintf(" title=\"%s\"", title)
	}
	if class != "" {
		link += fmt.Sprintf(" class=\"%s\"", class)
	}

	link += fmt.Sprintf(">%s</a>", text)

	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, link)
		return err
	})
}

func RenderImage(field interface{}) templ.Component {
	fmt.Println("RenderImage")
	fmt.Println(field)
	scField, ok := field.(model.ScField)
	if !ok {
		fmt.Println("RenderImage: not a ScField")
		return nil
	}
	fmt.Println(reflect.TypeOf(scField.Value))
	imageField, ok := scField.Value.(map[string]interface{})
	fmt.Println(reflect.TypeOf(imageField))
	if !ok {
		fmt.Println("RenderImage: not an imageField")
		return nil
	}

	src := imageField["src"].(string)
	alt := imageField["alt"].(string)
	width := imageField["width"].(string)
	height := imageField["height"].(string)

	img := fmt.Sprintf("<img src=\"%s\"", src)
	if alt != "" {
		img += fmt.Sprintf(" alt=\"%s\"", alt)
	}
	if width != "" {
		img += fmt.Sprintf(" width=\"%s\"", width)
	}
	if height != "" {
		img += fmt.Sprintf(" height=\"%s\"", height)
	}

	img += fmt.Sprintf(" />")

	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, img)
		return err
	})
}
