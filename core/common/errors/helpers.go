package errors

import (
	"strings"
)

func Unmarshal(data string) error {
	errors := strings.Split(data, "\n")
	var finalError error

	for i := len(errors) - 1; i >= 0; i-- {
		e := errors[i]
		var err Error
		se := strings.Split(e, " ")

		switch len(se) {
		case 2:
			err.Location = se[0]
			err.Msg = se[1]
		case 3:
			err.Location = se[0]
			if isCode(se[1]) {
				err.Code = extractCode(se[1])
				err.Msg = se[2]
			} else {
				err.Msg = strings.Join(se[1:], " ")
			}
		default:
			continue
		}
		finalError = Wrap(finalError, &err)
	}
	return finalError
}

func extractCode(code string) string {
	return strings.TrimRight(code, ":")
}

func isCode(code string) bool {
	// ascii code for ":" is 58
	return code[len(code)-1] == 58
}

// Top since errors can be wrapped and stacked,
// it's necessary to get the top level error for tests and validations
func Top(err error) string {
	switch t := err.(type) {
	case *Error:
		return t.top()
	case *withError:
		switch ct := t.current.(type) {
		case *Error:
			return ct.top()
		case *withError:
			return Top(t.current)
		default:
			return err.Error()
		}
	default:
		return err.Error()
	}
}

// PPrint - pretty print
func PPrint(err error) string {
	if err == nil {
		return ""
	}
	switch e := err.(type) {
	case *Error:
		return e.pprint()
	case *withError:
		return e.pprint()
	default:
		return e.Error()
	}
}

// isError - parses the error into Error
func isError(err error) *Error {
	t, ok := err.(*Error)
	if ok {
		return t
	}
	return nil
}

// isError - parses the error into withError
func isWithError(err error) *withError {
	t, ok := err.(*withError)
	if ok {
		return t
	}
	return nil
}
