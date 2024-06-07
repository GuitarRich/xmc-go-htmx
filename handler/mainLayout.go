package handler

import (
	"fmt"

	"github.com/guitarrich/headless-go-htmx/model"
	"github.com/guitarrich/headless-go-htmx/view/layout"
	"github.com/labstack/echo/v4"
)

type MainLayoutHandler struct{}

func (h MainLayoutHandler) HandleLayout(c echo.Context) error {

	fmt.Println("MainLayoutHandler")
	fmt.Println(c.Request().URL.Path)
	pageData := model.PageData{
		Title: "Headless Go [" + c.Request().URL.Path + "]",
		Body:  "Hello World",
	}

	return render(c, layout.MainLayout(pageData))
}
