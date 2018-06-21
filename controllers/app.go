package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

// APIAppPing is a healthcheck
func APIAppPing(context echo.Context) error {
	return context.NoContent(http.StatusOK)
}
