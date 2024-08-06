package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()

	// Route to serve the main page
	e.GET("/", helloHandler)

	// Route to handle the button click
	e.GET("/update", updateHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

// Handler for the main page
func helloHandler(c echo.Context) error {
	html := `
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Hello htmx</title>
        <script src="https://unpkg.com/htmx.org@2.0.1"></script>
    </head>
    <body>
        <h1>Hello, htmx!</h1>
        <button hx-get="/update" hx-swap="outerHTML">Click me</button>
    </body>
    </html>`
	return c.HTML(http.StatusOK, html)
}

// Handler for the update button click
func updateHandler(c echo.Context) error {
	return c.HTML(http.StatusOK, "<h1>Hello, htmx! Updated!</h1>")
}
