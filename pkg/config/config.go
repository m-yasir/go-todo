package config

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const LOGGER_FORMAT = "proto=${protocol}, method=${method}, uri=${uri}, status=${status}\n"
const DEFAULT_TIMEOUT = 5000

func ConfigureMiddlewares(app *echo.Echo) {
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: LOGGER_FORMAT,
	}))
	app.Use(middleware.CORS())
	app.Use(middleware.Gzip())
	app.Use(middleware.Recover())
}
