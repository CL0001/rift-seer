package renderer

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

type Templates struct {
	templates *template.Template
}

func NewRenderer() *Templates {
	tmpl := template.Must(template.ParseGlob("views/*.html"))
	return &Templates{templates: tmpl}
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}