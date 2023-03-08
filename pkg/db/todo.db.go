package db

import (
	m "github.com/m-yasir/go-todo/pkg/models"
)

var Todos []m.Todo

func SeedTodosDb() {
	Todos = []m.Todo{
		{
			Id: "1", Name: "First task", IsCompleted: false,
		},
		{
			Id: "2", Name: "Second task", IsCompleted: true,
		},
		{
			Id: "3", Name: "Third task", IsCompleted: false,
		},
		{
			Id: "4", Name: "Test task", IsCompleted: true,
		},
	}
}
