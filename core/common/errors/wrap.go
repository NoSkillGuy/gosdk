package errors

import (
	"strings"
)

// Wrap wrap the previous error with current error/ message
func Wrap(previous error, current interface{}) error {
	var currentError error
	switch c := current.(type) {
	case error:
		// fmt.Println("------------------- 1")
		if c == nil {
			// fmt.Println("------------------- 1.1")
			currentError = invalidWrap()
		} else {
			// fmt.Println("------------------- 1.2")
			currentError = c
		}
	case string:
		// fmt.Println("------------------- 2")
		if strings.TrimSpace(c) == "" {
			currentError = invalidWrap()
		} else {
			currentError = newWithLevel(3, c)
		}
	default:
		// fmt.Println("------------------- 3")
		currentError = invalidWrap()
	}

	return &withError{
		previous: previous,
		current:  currentError,
	}
}

func invalidWrap() *Error {
	code := "incorrect_usage"
	msg := "you should pass either error or message to properly wrap the error!"
	return newWithLevel(4, code, msg)
}
