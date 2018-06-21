package main

import (
	"github.com/antonve/api-boilerplate/migrations"
	"github.com/antonve/api-boilerplate/utils"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"log"
)

func main() {
	migrations.Migrate()

	// Echo instance
	e := echo.New()
	log.Println("Starting boilerplate API")

	// Middleware
	e.Use(middleware.Recover())
	defer utils.SetupErrorLogging(e)()

	// Routes
	utils.SetupRouting(e)

	// Start server
	e.Start(":7000")
}
