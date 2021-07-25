package errors

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type wrapTopTestCase struct {
	about              string
	testCase           []interface{}
	expectedTopMessage string
}

func getWrapTopTestCases() []wrapTopTestCase {
	return []wrapTopTestCase{
		// {
		// 	about: "wrapping all errors",
		// 	testCase: []interface{}{
		// 		New("500", "This is a very big error! Beware of it!"),
		// 		New("", "This is a very big error! Beware of it!"),
		// 		New("401", ""),
		// 		New("This is a short error!"),
		// 		New("code", "message", "third"),
		// 		New("code", "message", "third", "fourth"),
		// 		New(),
		// 		errors.New("error created from err package"),
		// 		fmt.Errorf("%s", "error created from fmt package"),
		// 		nil,
		// 	},
		// 	expectedTopMessage: "incorrect_usage: you should pass either error or message to properly wrap the error!",
		// },
		// {
		// 	about: "wrapping all messages",
		// 	testCase: []interface{}{
		// 		"This is a very \"big\" error! Beware of it!",
		// 		"This is a very 'big' error! Beware of it!",
		// 		"This is a short error!",
		// 		"",
		// 	},
		// 	expectedTopMessage: "incorrect_usage: you should pass either error or message to properly wrap the error!",
		// },
		// {
		// 	about: "wrapping errors and messages",
		// 	testCase: []interface{}{
		// 		New("500", "This is a very big error! Beware of it!"),
		// 		"This is a very \"big\" error! Beware of it!",
		// 		New("401", ""),
		// 		"This is a very 'big' error! Beware of it!",
		// 		New("This is a short error!"),
		// 		"",
		// 		nil,
		// 		New("code", "message", "third"),
		// 		"This is a short error!",
		// 		New("code", "message", "third", "fourth"),
		// 		New(),
		// 		New("", "This is a very big error! Beware of it!"),
		// 	},
		// 	expectedTopMessage: "This is a very big error! Beware of it!",
		// },
		{
			about: "wrapping error with nil error",
			testCase: []interface{}{
				New("500", "This is a very big error! Beware of it!"),
				nil,
				errors.New(""),
			},
			expectedTopMessage: "incorrect_usage: you should pass either error or message to properly wrap the error!",
		},
	}
}

func TestWrap(t *testing.T) {
	for _, gtc := range getWrapTopTestCases() {
		t.Run(gtc.about, func(t *testing.T) {
			var wrappedError error
			for _, tc := range gtc.testCase {
				wrappedError = Wrap(wrappedError, tc)
			}

			require.Equal(t, len(gtc.testCase), len(strings.Split(wrappedError.Error(), "\n")))
		})
	}
}
