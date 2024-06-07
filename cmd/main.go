package main

import (
	"github.com/guitarrich/headless-go-htmx/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()

	app.Static("/css", "public/css")
	app.Static("/js", "public/js")
	app.Static("/img", "public/img")

	app.Use(middleware.Logger())

	layoutHandler := handler.MainLayoutHandler{}
	app.GET("/*", layoutHandler.HandleLayout)

	app.Start(":42069")
}
