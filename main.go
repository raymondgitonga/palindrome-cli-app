package main

import (
	"fmt"
	"github.com/raymondgitonga/palindrome-cli-app/cmd"
)

func main() {
	err := cmd.RunPalindromeCommand("word")

	if err != nil {
		fmt.Println(err)
	}

}
