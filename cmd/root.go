package cmd

import (
	"fmt"
	"github.com/raymondgitonga/palindrome-cli-app/internal"
	"github.com/spf13/cobra"
)

func RunPalindromeCommand(flag string) *cobra.Command {
	command := &cobra.Command{
		Use:   "palindrome",
		Short: "Checks is a a word entered is a palindrome",
		Long:  `Checks is a a word entered is a palindrome`,
		RunE: func(cmd *cobra.Command, args []string) error {
			word, err := cmd.Flags().GetString(flag)

			if err != nil {
				return fmt.Errorf("Error reading word flag: %v ", err)
			}

			if len(word) <= 0 {
				return fmt.Errorf("Kindly add flag --word. Example  go run main.go palindrome --word={word you want to check}")
			}

			result := checkIsPalindrome(word)

			fmt.Println(result)
			return nil
		},
	}
	command.PersistentFlags().String(flag, "", "Check if word is a palindrome")

	return command
}

func checkIsPalindrome(word string) string {
	if internal.IsPalindrome(word) {
		return "This is for sure a palindrome."
	}
	return "This is not a palindrome."
}
