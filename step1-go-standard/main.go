package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
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
	fmt.Fprint(w, html)
}
