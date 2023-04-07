package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var Server *echo.Echo

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("static/*.html")),
	}

	// Echo server Instance
	Server = echo.New()

	// Add template render engine
	Server.Renderer = t

	// Middlewares
	Server.Use(middleware.Logger())
	Server.Use(middleware.Recover())

	// Serve static files
	Server.Static("/static", "./static")

	Server.GET("/", Index)
	Server.POST("/rooms", EnterChatRoom)
	Server.GET("/ws/rooms/:id", ChatHandler)

	Server.Logger.Fatal(Server.Start(":1323"))
}
