package cmd_test

import (
	"fmt"
	"github.com/raymondgitonga/palindrome-cli-app/cmd"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRootCmd(t *testing.T) {
	var result string
	type testCase struct {
		input  string
		result string
		err    error
	}

	testCases := []testCase{
		{
			input:  "word",
			result: "This is not a palindrome.",
			err:    nil,
		},
		{
			input:  "racecar",
			result: "This is for sure a palindrome.",
			err:    nil,
		},
	}

	for _, tc := range testCases {
		palindromeCmd := &cobra.Command{
			Use: "palindrome",
			RunE: func(command *cobra.Command, args []string) error {
				if len(tc.input) <= 0 {
					return fmt.Errorf("Kindly add flag --word. Example  go run main.go palindrome --word={word you want to check}")
				}

				result = cmd.CheckIsPalindrome(tc.input)
				return nil
			},
		}

		err := palindromeCmd.Execute()

		assert.Equal(t, tc.err, err)
		assert.Equal(t, tc.result, result)
	}
}
