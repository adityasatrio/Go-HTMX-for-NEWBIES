# Go HTMX for NEWBIES

This project is a step-by-step guide to building web applications using Go, Echo framework, and HTMX. It's designed for beginners who want to learn go + htmx.

## Prerequisites

Before starting this tutorial, make sure you have:

- Go installed on your system (version 1.16 or later recommended)
- Basic understanding of HTML and JavaScript
- A text editor or IDE of your choice

## Project Structure

```
├── Readme.md
├── go.mod
├── go.sum
├── step1-go-standard
├── step2-go-standard-update-handler
├── step3-go-echo
├── step4-go-echo-template
├── step5-go-echo-template-html-component
└── step6-go-echo-template-htmx-form
```

## Step-by-Step Guide

### Step 1: Go Standard

**Folder**: `step1-go-standard`

This step introduces the basics of creating a web server using Go's standard library. You'll learn how to:
- Set up a simple HTTP server
- Show HTML output in response

### Step 2: Go Standard Update Handler

**Folder**: `step2-go-standard-update-handler`

Building on Step 1, this section focuses on:
- Working with hard-coded HTML in Go
- Responding with the whole HTML
- Calling an update handler using HTMX `hx-get`

### Step 3: Go Echo

**Folder**: `step3-go-echo`

This step introduces the Echo framework. You'll learn:
- Setting up a simple Echo server
- Basic routing with Echo
- Hard-coding HTML and calling API update handler with `hx-get`

### Step 4: Go Echo Template

**Folder**: `step4-go-echo-template`

Here, we dive into using HTML templates with Echo. This section covers:
- Setting up the template renderer
- Creating basic HTML templates
- Passing data from your Go code to the templates

### Step 5: Go Echo Template HTML Component

**Folder**: `step5-go-echo-template-html-component`

This step focuses on modularizing your HTML using components. You'll learn:
- How to break down your HTML into reusable components

### Step 6: Go Echo Template HTMX Form

**Folder**: `step6-go-echo-template-htmx-form`

The final step introduces more advanced HTMX usage. This section covers:
- Integrating HTMX with your Echo application
- Creating dynamic forms with HTMX `hx-post`
- Handling file uploads

## Additional Resources

- `adminlte-dist`: This folder contains the AdminLTE template files, which are used to create a professional-looking UI in Step 6.
- `go.mod` and `go.sum`: These files manage the project's dependencies.

## Getting Started

To get started with this project:

1. Clone the repository
2. Navigate to each step's folder
3. Run the main.go file in each step using `go run main.go`
4. Open your browser and visit `http://localhost:8080` (or the port specified in the code)

## Further Learning

- [Official Go Documentation](https://golang.org/doc/)
- [Echo Framework Documentation](https://echo.labstack.com/)
- [HTMX Documentation](https://htmx.org/docs/)