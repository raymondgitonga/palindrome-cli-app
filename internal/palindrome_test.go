package internal_test

import (
	"github.com/raymondgitonga/palindrome-cli-app/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IsPalindrome(t *testing.T) {
	testCases := []struct {
		input  string
		result bool
	}{
		{
			input:  "john",
			result: false,
		},
		{
			input:  "racecar",
			result: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			isPalindrome := internal.IsPalindrome(tc.input)
			assert.Equal(t, tc.result, isPalindrome)
		})
	}

}
