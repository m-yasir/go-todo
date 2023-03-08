package utils

import (
	"net/url"
	"testing"
)

// TestStrToBoolEmpty calls utils.StrToBool with a stringified bool value,
// for a valid return value and no errors.
func TestStrToBoolOnTrue(t *testing.T) {
	want := true
	val, err := StrToBool("true")
	if err != nil {
		t.Fatalf(`StrToBool("true") resulted in an error reason: %q, val = %v, want match for %#v`, err, val, want)
	}
	if want != val {
		t.Fatalf(`StrToBool("Gladys") = %v, %v, want match for %#v, nil`, val, err, want)
	}
}

// TestStrToBoolEmpty calls utils.StrToBool with an empty string,
// checking for an error.
func TestStrToBoolEmpty(t *testing.T) {
	val, err := StrToBool("")
	want := false
	if err != nil {
		t.Fatalf(`StrToBool("") = %v, want %v, error %v`, val, want, err)
	}
	if want != val {
		t.Fatalf(`StrToBool("") = %v, want %v, error %q`, val, want, err)
	}
}

type TestParserStruct struct {
	msg string
}

func (tp TestParserStruct) ParseParams(params url.Values) TestParserStruct {
	tp.msg = params.Get("test")
	return tp
}

func TestParseQueryParam(t *testing.T) {

	val := ParseQueryParam[TestParserStruct](url.Values{"test": {"hello"}}, TestParserStruct{})

	want := "hello"

	if val.msg != want {
		t.Fatalf(`ParseQueryParam(url.Values, TestParserStruct{}) = %q, want "hello"`, val)
	}
}
