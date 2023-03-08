package setup

import (
	"github.com/m-yasir/go-todo/pkg/db"
	routes "github.com/m-yasir/go-todo/pkg/routes"

	"github.com/labstack/echo/v4"
)

type RouteGroup struct {
	name    string
	handler func(*echo.Group)
}

func SetupGroupRoutes(app *echo.Echo, routeGroups []RouteGroup) {
	for _, routeGroup := range routeGroups {
		g := app.Group(routeGroup.name)
		routeGroup.handler(g)
	}
}

func SetupRoutes(app *echo.Echo) {
	// Setup individual routes
	app.GET("/ping", routes.HandlePing)

	// Setup grouped routes
	SetupGroupRoutes(app, []RouteGroup{
		{"/todos", routes.HandleTodoRoutes},
	})
}

func SetupDb() {
	db.SeedTodosDb()
}
