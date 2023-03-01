package setup

import (
	routes "github.com/m-yasir/go-todo/pkg/routes"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(app *echo.Echo) {
	app.GET("/health", routes.HandleHealth)
	app.GET("/ping", routes.HandlePing)
}
