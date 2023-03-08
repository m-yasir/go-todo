package controllers

import (
	"net/url"
	"strings"

	"github.com/m-yasir/go-todo/pkg/db"
	m "github.com/m-yasir/go-todo/pkg/models"
	"github.com/m-yasir/go-todo/pkg/utils"
)

func GetTodos(filters url.Values) []m.Todo {
	todos := make([]m.Todo, 0)

	// replace with db query instead
	for _, Todo := range db.Todos {
		contains := strings.ToLower(filters.Get("contains"))
		completed := utils.StrToBool(filters.Get("completed"))

		// check if current Todo complies with completed and contains filter
		shouldInclude := (filters.Get("completed") == "" || completed == Todo.IsCompleted) && (contains == "" || strings.Contains(strings.ToLower(Todo.Name), contains))

		// update todos db
		if shouldInclude {
			todos = append(todos, Todo)
		}
	}
	return todos
}

func CreateTodo(todo m.Todo) []m.Todo {
	db.Todos = append(db.Todos, todo)
	return db.Todos
}
