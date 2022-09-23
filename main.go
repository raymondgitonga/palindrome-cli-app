/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"github.com/raymondgitonga/palindrome-cli-app/cmd"
)

func main() {
	command := cmd.RunPalindromeCommand("word")
	command.PersistentFlags().String("word", "", "Check if word is a palindrome")
	err := command.Execute()

	if err != nil {
		fmt.Printf("Error running palindrome command %s\n", err)
	}

}
