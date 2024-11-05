package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/guitarrich/headless-go-htmx/middleware/editing"
	"github.com/guitarrich/headless-go-htmx/pipelines/request"
	"github.com/guitarrich/headless-go-htmx/sitecore"
	"github.com/guitarrich/headless-go-htmx/sitecore/schandlers"
	"github.com/guitarrich/headless-go-htmx/view/components"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/http2"
)

func main() {
	app := echo.New()

	app.Static("/favicon.ico", "favicon.ico")
	app.Static("/css", "css")
	app.Static("/js", "js")
	app.Static("/img", "img")

	app.Use(middleware.Logger())

	components.RegisterComponents()
	layoutHandler := request.RequestPipelineHandler{}
	app.GET("/*", layoutHandler.RequestBeginHandler)

	// Editing API
	editingHandler := editing.EditingRequestHandler{}
	app.GET("/api/editing/config", editingHandler.Config)
	app.GET("/api/editing/render", editingHandler.Render)

	app.GET("/-/media/*", schandlers.NewMediaHandler().ServeHTTP)

	s := &http2.Server{
		MaxConcurrentStreams: 250,
		MaxReadFrameSize:     1048576,
		IdleTimeout:          10 * time.Second,
	}

	port := sitecore.GetEnvVar("PORT")
	if err := app.StartH2CServer(fmt.Sprintf(":%s", port), s); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
