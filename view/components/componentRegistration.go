package components

import "github.com/guitarrich/headless-go-htmx/sitecore/render"

func RegisterComponents() {
	render.RegisterComponent("Header", Header)
	render.RegisterComponent("Hero", Hero)
	render.RegisterComponent("Container", Container)
	render.RegisterComponent("Image", Image)
	render.RegisterComponent("Promo", Promo)
	render.RegisterComponent("RichText", RichText)
	render.RegisterComponent("Navigation", Navigation)
}
