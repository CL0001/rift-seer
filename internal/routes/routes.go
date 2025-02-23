package routes

import (
	"github.com/CL0001/rift-seer/internal/auth"
	"github.com/CL0001/rift-seer/internal/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(router *echo.Echo) {
	router.GET("/", handlers.HomePage)
	router.GET("/auth", handlers.AuthPage)
	router.GET("/about", handlers.AboutPage)

	router.POST("/auth/register", auth.RegisterUser)
	router.POST("/auth/login", auth.LoginUser)

	router.GET("/dashboard", handlers.DashboardPage)

//	router.GET("/dashboard", auth.IsAuthenticated())
}
