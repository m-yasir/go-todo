package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/m-yasir/go-todo/pkg/controllers"
	"github.com/m-yasir/go-todo/pkg/models"
	"github.com/m-yasir/go-todo/pkg/utils"
)

func createTodo(c echo.Context) error {
	return c.JSON(http.StatusOK, `{ "data": [], "success": true }`)
}

func updateTodo(c echo.Context) error {
	return c.JSON(http.StatusOK, `{ "data": [], "success": true }`)
}

func deleteTodo(c echo.Context) error {
	return c.JSON(http.StatusOK, `{ "data": [], "success": true }`)
}

func getTodos(c echo.Context) error {
	todos := controllers.GetTodos(c.QueryParams())
	return c.JSON(http.StatusOK, utils.Response[models.Todo]{Data: todos, Success: true})
}

func HandleTodoRoutes(g *echo.Group) {
	g.GET("", getTodos)
	g.POST("", createTodo)
	g.PATCH("/:id", updateTodo)
	g.DELETE("/:id", deleteTodo)
}
