package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HomePage(context echo.Context) error {
	data := make(map[string]string)

	return context.Render(http.StatusOK, "index.html", data)
}

func AuthPage(context echo.Context) error {
	data := make(map[string]string)

	return context.Render(http.StatusOK, "auth.html", data)
}

func AboutPage(context echo.Context) error {
	data := make(map[string]string)

	return context.Render(http.StatusOK, "about.html", data)
}

func DashboardPage(context echo.Context) error {
	data := make(map[string]string)

	return context.Render(http.StatusOK, "dashboard.html", data)
}