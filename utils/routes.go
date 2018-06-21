package utils

import (
	"github.com/antonve/api-boilerplate/controllers"
	"github.com/labstack/echo"
)

// SetupRouting Define all routes here
func SetupRouting(e *echo.Echo) {
	// Routes
	routesAPI := e.Group("/api")
	routesAPI.GET("/ping", echo.HandlerFunc(controllers.APIAppPing))
}
