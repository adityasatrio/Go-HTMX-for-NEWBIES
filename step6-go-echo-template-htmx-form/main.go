package main

import (
	"bufio"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
		templates: template.Must(template.ParseGlob("step6-go-echo-template-htmx-form/adminlte-dist/**/*.html")),
	}
	e.Renderer = renderer
	e.Static("/static", "step6-go-echo-template-htmx-form/adminlte-dist")

	// Route to serve the main page
	e.GET("/", helloHandler)

	// Route to handle the button click
	e.POST("/upload", uploadHandler)
	e.POST("/upload-read", uploadReadHandler)
	e.POST("/upload-read-htmx", uploadReadHtmxHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

// Handler for the main page
func helloHandler(c echo.Context) error {
	data := map[string]interface{}{
		"Title":       "Hello htmx",
		"HeaderTitle": "Welcome to My Site",
		"BodyTitle":   "Data Table",
		"Items": []map[string]interface{}{
			{"ID": 1, "Name": "Item 1"},
			{"ID": 2, "Name": "Item 2"},
		},
	}
	return c.Render(http.StatusOK, "form-file.html", data)
}

// Handler for file upload
func uploadHandler(c echo.Context) error {
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(filepath.Join("step6-go-echo-template-htmx-form/uploads", file.Filename))
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	//return c.String(http.StatusOK, "File uploaded successfully: "+file.Filename)

	data := map[string]interface{}{
		"Title":       "Hello htmx",
		"HeaderTitle": "Welcome to My Site",
		"BodyTitle":   "Data Table",
		"Items": []map[string]interface{}{
			{"ID": 1, "Name": "Item 1"},
			{"ID": 2, "Name": "Item 2"},
		},
		"fileName": "File Uploaded Name is " + file.Filename,
	}

	return c.Render(http.StatusOK, "form-file.html", data)
}

// Handler for file upload and read
func uploadReadHandler(c echo.Context) error {
	// Source
	file, err := c.FormFile("file-read")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Read file content
	reader := bufio.NewReader(src)
	var content string
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		content += line + "<br>"
	}

	// Return content
	data := map[string]interface{}{
		"Title":       "File Upload and Read",
		"HeaderTitle": "Welcome to My Site",
		"FileContent": content,
		"fileName":    "File Uploaded Name is " + file.Filename,
	}
	return c.Render(http.StatusOK, "form-file.html", data)

}

// Handler for file upload and read
func uploadReadHtmxHandler(c echo.Context) error {
	// Source
	file, err := c.FormFile("file-read-htmx")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Read file content
	reader := bufio.NewReader(src)
	var content string
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		content += line + "<br>"
	}

	// Return content
	//return c.Render(http.StatusOK, "form-file.html", data)
	return c.HTML(http.StatusOK, content)

}
