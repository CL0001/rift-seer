package main

import (
	"github.com/CL0001/rift-seer/internal/renderer"
	"github.com/CL0001/rift-seer/internal/routes"
	"github.com/CL0001/rift-seer/pkg/db"
	"github.com/CL0001/rift-seer/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	utils.LoadEnv()

	db.InitDB()

	app := echo.New()
	app.Use(middleware.Logger())

	app.Renderer = renderer.NewRenderer()

	routes.RegisterRoutes(app)

	app.Logger.Fatal(app.Start(":8000"))
}