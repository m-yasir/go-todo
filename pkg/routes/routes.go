package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandlePing(c echo.Context) error {
	return c.String(http.StatusOK, "Pong")
}

func HandleHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, `{ "alive": true }`)
}
