package routes

import (
	"github.com/CL0001/rift-seer/internal/auth"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(router *echo.Echo) {
	router.Static("/", "web/static")

	router.File("/", "index.html")

	router.POST("/register", auth.RegisterUser)
	router.POST("/login", auth.LoginUser)

	router.GET("/dashboard", auth.IsAuthenticated(test))
}
