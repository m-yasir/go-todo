package utils

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrToBool(t *testing.T) {
	t.Run("converts stringified boolean to bool", func(t *testing.T) {
		val := StrToBool("true")
		want := true

		assert.True(t, val == want)
	})
	t.Run("returns false if empty string is passed", func(t *testing.T) {
		val := StrToBool("")
		want := false

		assert.True(t, val == want)
	})
}

type TestParserStruct struct {
	msg string
}

func (tp TestParserStruct) ParseParams(params url.Values) TestParserStruct {
	tp.msg = params.Get("test")
	return tp
}

func TestParseQueryParam(t *testing.T) {

	t.Run("Parses query parameters correctly", func(t *testing.T) {
		val := ParseQueryParam[TestParserStruct](url.Values{"test": {"hello"}}, TestParserStruct{})

		want := "hello"

		if val.msg != want {
			t.Fatalf(`ParseQueryParam(url.Values, TestParserStruct{}) = %q, want "hello"`, val)
		}
	})
}
