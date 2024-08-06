package main

import (
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Instantiate a template renderer
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("step4-go-echo-template/templates/*.html")),
	}
	e.Renderer = renderer

	// Route to serve the main page
	e.GET("/", helloHandler)

	// Route to handle the button click
	e.GET("/update", updateHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

// Handler for the main page
func helloHandler(c echo.Context) error {
	data := map[string]interface{}{
		"Name": "htmx jhoni",
	}

	return c.Render(http.StatusOK, "index.html", data)
}

// Handler for the update button click
func updateHandler(c echo.Context) error {
	data := map[string]interface{}{
		"Name": "htmx aditya",
	}

	return c.Render(http.StatusOK, "update.html", data)
}
