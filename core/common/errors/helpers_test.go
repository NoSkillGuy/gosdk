package errors

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTop(t *testing.T) {
	for _, gtc := range getWrapTopTestCases() {
		t.Run(gtc.about, func(t *testing.T) {
			var wrappedError error
			for _, tc := range gtc.testCase {
				wrappedError = Wrap(wrappedError, tc)
			}
			require.Equal(t, gtc.expectedTopMessage, Top(wrappedError))
		})
	}
}
