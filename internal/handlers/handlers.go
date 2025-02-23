package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HomePage(context echo.Context) error {
	data := make(map[string]string)

	return context.Render(http.StatusOK, "index.html", data)
}