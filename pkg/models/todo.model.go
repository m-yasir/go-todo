package models

import (
	"net/url"

	"github.com/m-yasir/go-todo/pkg/utils"
)

type Todo struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	IsCompleted bool   `json:"isCompleted"`
}

type TodoFilters struct {
	Completed bool
	Contains  string
}

func (t TodoFilters) ParseParams(params url.Values) TodoFilters {
	if params.Has("contains") {
		t.Contains = params.Get("contains")
	}
	if params.Has("Completed") {
		t.Completed = utils.StrToBool(params.Get("completed"))
	}
	return t
}
