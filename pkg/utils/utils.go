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

func StrToBool(str string) (bool, error) {
	result, err := strconv.ParseBool(str)
	if err != nil {
		result = false
	}
	return result, nil
}
