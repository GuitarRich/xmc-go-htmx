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
	"github.com/guitarrich/headless-go-htmx/view/components/atoms"
	"strings"
)

func buildStyle(props model.PlaceholderComponent) string {
	imageUrl := render.GetImageField(props.Fields, "HeroImage").Value.Src
	if strings.Contains(imageUrl, "/-/media/") {
		trimTo := strings.Index(imageUrl, "/-/media/")
		imageUrl = imageUrl[trimTo : len(imageUrl)-1]
	}
	return fmt.Sprintf("--image-url:url('%s');", imageUrl)
}

const backgroundStyle = "component hero bg-white bg-[image:var(--image-url)] bg-cover bg-center"

func Hero(component model.PlaceholderComponent, sc model.SitecoreContext) templ.Component {
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
		templ_7745c5c3_Err = templ.Raw(
			fmt.Sprintf("<div style=\"%s\" class=\"%s\">", buildStyle(component),
				render.DecorateComponent(backgroundStyle, component)),
		).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"relative overflow-hidden py-24 lg:py-32\"><div aria-hidden=\"true\" class=\"flex absolute -top-96 start-1/2 transform -translate-x-1/2\"><div class=\"bg-gradient-to-r opacity-50 blur-3xl w-[90rem] h-[70rem] rounded-full origin-top-left -rotate-12 -translate-x-[15rem] from-transparent via-scblack to-transparent\"></div></div><div class=\"relative z-10\"><div class=\"container max-w-full w-full py-10 lg:py-16\"><div class=\"max-w-2xl text-center mx-auto text-scwhite\"><div class=\"\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = render.RichTextField(component.Fields, "HeroTagLine").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"mt-5 max-w-2xl\"><h1 class=\"scroll-m-20 text-5xl font-normal tracking-wider lg:text-5xl\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = render.RichTextField(component.Fields, "HeroTitle").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h1></div><div class=\"mt-5 max-w-3xl\"><div class=\"text-xl text-scgray-90\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = render.RichTextField(component.Fields, "HeroText").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><div class=\"mt-8 gap-3 flex justify-center\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = atoms.Button(atoms.ButtonProps{
			Size: atoms.ButtonSizeLarge,
		}, component.Fields, "HeroLink").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.Raw(`
</div>`).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
