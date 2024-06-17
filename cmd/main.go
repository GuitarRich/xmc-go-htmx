package main

import (
	"github.com/guitarrich/headless-go-htmx/handler"
	"github.com/guitarrich/headless-go-htmx/view/components"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()

	app.Static("/favicon.ico", "favicon.ico")
	app.Static("/css", "css")
	app.Static("/js", "js")
	app.Static("/img", "img")

	app.Use(middleware.Logger())

	components.RegisterComponents()
	layoutHandler := handler.MainLayoutHandler{}
	app.GET("/*", layoutHandler.HandleLayout)

	app.Start(":42069")
}
