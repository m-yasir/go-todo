package main

import (
	"fmt"

	config "github.com/m-yasir/go-todo/pkg/config"
	setup "github.com/m-yasir/go-todo/pkg/setup"

	"github.com/labstack/echo/v4"
)

func init() {
	// setup db
}

var serverPort string = ":8080"

func main() {

	app := echo.New()

	config.ConfigureMiddlewares(app)
	setup.SetupRoutes(app)

	if err := app.Start(serverPort); err != nil {
		fmt.Printf("Unable to start server %s\n", err)
	}
}
