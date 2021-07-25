package errors

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type newErrorTestCase struct {
	about           string
	args            []string
	expectedCode    string
	expectedMessage string
}

func getNewErrorTestCase() []newErrorTestCase {
	return []newErrorTestCase{
		{
			about:           "creating an error with code and msg.",
			args:            []string{"500", "This is a very big error! Beware of it!"},
			expectedCode:    "500",
			expectedMessage: "This is a very big error! Beware of it!",
		},
		{
			about:           "creating an error with empty code and msg.",
			args:            []string{"", "This is a very big error! Beware of it!"},
			expectedCode:    "",
			expectedMessage: "This is a very big error! Beware of it!",
		},
		{
			about:           "creating an error with code and empty msg.",
			args:            []string{"401", ""},
			expectedCode:    "401",
			expectedMessage: "",
		},
		{
			about:           "creating an error with just msg.",
			args:            []string{"This is a short error!"},
			expectedCode:    "",
			expectedMessage: "This is a short error!",
		},
		{
			about:           "creating an error by passing 3 parameters which is not allowed",
			args:            []string{"code", "message", "third"},
			expectedCode:    "incorrect_usage",
			expectedMessage: "max allowed parameters is 2 i.e code, msg. parameters sent - 3",
		},
		{
			about:           "creating an error by passing 4 parameters which is not allowed",
			args:            []string{"code", "message", "third", "fourth"},
			expectedCode:    "incorrect_usage",
			expectedMessage: "max allowed parameters is 2 i.e code, msg. parameters sent - 4",
		},
		{
			about:           "creating an error with empty parameters",
			args:            []string{},
			expectedCode:    "incorrect_usage",
			expectedMessage: "max allowed parameters is 2 i.e code, msg. parameters sent - 0",
		},
		{
			about:           "creating an error with spaces in code",
			args:            []string{"This is a very long code", "This is its very long message"},
			expectedCode:    "incorrect_code",
			expectedMessage: "code should not have spaces. use 'this_is_a_very_long_code' instead of 'This is a very long code'",
		},
	}
}

func TestNew(t *testing.T) {
	for _, tc := range getNewErrorTestCase() {
		t.Run(tc.about, func(t *testing.T) {
			err := New(tc.args...)

			require.Equal(t, tc.expectedCode, err.Code)
			require.Equal(t, tc.expectedMessage, err.Msg)
		})
	}
}

func TestError(t *testing.T) {
	for _, tc := range getNewErrorTestCase() {
		t.Run(tc.about, func(t *testing.T) {
			err := New(tc.args...)

			require.Contains(t, err.Error(), tc.expectedMessage)
		})
	}
}

type newErrorfTestCase struct {
	about           string
	code            string
	format          string
	args            []interface{}
	expectedCode    string
	expectedMessage string
}

func getNewErrorfTestCase() []newErrorfTestCase {
	return []newErrorfTestCase{
		{
			about:           "creating an error with code and msg with integer arg.",
			code:            "integer_error",
			format:          "This error has a integer: %d",
			args:            []interface{}{500},
			expectedCode:    "integer_error",
			expectedMessage: "This error has a integer: 500",
		},
		{
			about:           "creating an error with code and msg with string arg.",
			code:            "string_error",
			format:          "This error has a string: %s",
			args:            []interface{}{"500"},
			expectedCode:    "string_error",
			expectedMessage: "This error has a string: 500",
		},
		{
			about:           "creating an error with empty code and empty msg with string arg.",
			code:            "",
			format:          "This error has empty code with a string: %s",
			args:            []interface{}{"401"},
			expectedCode:    "",
			expectedMessage: "This error has empty code with a string: 401",
		},
		{
			about:           "creating an error with just msg.",
			code:            "",
			format:          "This is a short error!",
			args:            []interface{}{},
			expectedCode:    "",
			expectedMessage: "This is a short error!",
		},
		{
			about:           "creating an error with code and format which expects values but not sending values",
			code:            "code",
			format:          "This format expects integer value %d",
			args:            []interface{}{},
			expectedCode:    "code",
			expectedMessage: "This format expects integer value %!d(MISSING)",
		},
		{
			about:           "creating an error with code and format which expects integer value but we are sending string value",
			code:            "code",
			format:          "This format expects integer value %d",
			args:            []interface{}{"string value"},
			expectedCode:    "code",
			expectedMessage: "This format expects integer value %!d(string=string value)",
		},
		{
			about:           "creating an error with format and values but having empty code",
			code:            "",
			format:          "This is a sample format having %s",
			args:            []interface{}{"string value"},
			expectedCode:    "",
			expectedMessage: "This is a sample format having string value",
		},
		{
			about:           "creating an error with just format",
			code:            "",
			format:          "The format error",
			args:            []interface{}{},
			expectedCode:    "",
			expectedMessage: "The format error",
		},
		{
			about:           "creating an error with no code, format but have values",
			code:            "",
			format:          "",
			args:            []interface{}{"arg1", 2, "arg3"},
			expectedCode:    "",
			expectedMessage: "%!(EXTRA string=arg1, int=2, string=arg3)",
		},
		{
			about:           "creating an error with extra values",
			code:            "code",
			format:          "This is a sample format",
			args:            []interface{}{"extra arg1", 2, "extra arg3", 3.445, nil},
			expectedCode:    "code",
			expectedMessage: "This is a sample format%!(EXTRA string=extra arg1, int=2, string=extra arg3, float64=3.445, <nil>)",
		},
	}
}

func TestNewf(t *testing.T) {
	for _, tc := range getNewErrorfTestCase() {
		t.Run(tc.about, func(t *testing.T) {
			err := Newf(tc.code, tc.format, tc.args...)

			require.Equal(t, tc.expectedCode, err.Code)
			require.Equal(t, tc.expectedMessage, err.Msg)
		})

	}
}
