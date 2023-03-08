package utils

import (
	"net/url"
	"strconv"
)

type QueryParamParser[T comparable] interface {
	ParseParams(url.Values) T
}

func ParseQueryParam[T comparable](params url.Values, parser QueryParamParser[T]) T {
	return parser.ParseParams(params)
}

func StrToBool(str string) bool {
	result, err := strconv.ParseBool(str)
	if err != nil {
		result = false
	}
	return result
}

type Response[T comparable] struct {
	Data    []T  `json:"data"`
	Success bool `json:"success"`
}
