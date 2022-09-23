package cmd

import (
	"fmt"
	"github.com/raymondgitonga/palindrome-cli-app/internal"
	"os"

	"github.com/spf13/cobra"
)

// palindromeCmd represents the palindrome command
var palindromeCmd = &cobra.Command{
	Use:   "palindrome",
	Short: "Checks is a a word entered is a palindrome",
	Long:  `Checks is a a word entered is a palindrome`,
	Run: func(cmd *cobra.Command, args []string) {
		word, err := cmd.Flags().GetString("word")

		if err != nil {
			fmt.Printf("Error reading word flag: %v ", err)
			return
		}

		if len(word) <= 0 {
			fmt.Printf("Kindly add flag --word. Example  go run main.go palindrome --word={word you want to check}")
			return
		}

		result := checkIsPalindrome(word)

		fmt.Println(result)
	},
}

func init() {
	palindromeCmd.PersistentFlags().String("word", "", "Check if word is a palindrome")
}

func Execute() {
	err := palindromeCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func checkIsPalindrome(word string) string {
	if internal.IsPalindrome(word) {
		return "This is for sure a palindrome."
	}
	return "This is not a palindrome."
}
