package controllers

import (
	"log"
	"runtime/debug"

	"github.com/antonve/api-boilerplate/config"
	"github.com/labstack/echo"
)

// Serve a request with errors
func ServeWithError(context echo.Context, statusCode int, err error) error {
	handleError(err)

	return context.NoContent(statusCode)
}

func handleError(err error) {
	log.Println(err.Error())

	if config.GetConfig().Debug {
		debug.PrintStack()
	}
}
