package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IsPalindrome(t *testing.T) {
	type testCase struct {
		input  string
		result bool
	}

	testCases := []testCase{
		{
			input:  "john",
			result: false,
		},
		{
			input:  "racecar",
			result: true,
		},
		{
			input:  "",
			result: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			isPalindrome := IsPalindrome(tc.input)
			assert.Equal(t, tc.result, isPalindrome)
		})
	}

}
