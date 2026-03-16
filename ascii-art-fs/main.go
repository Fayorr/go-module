package main

import (
	"fmt"
	"os"
)

func validator(arguments []string) bool {
	// Check for 1 or 2 arguments
	if len(arguments) > 2 || len(arguments) < 1 {
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		return false
	}
	if arguments[0] == "" {
		return false
	}
	if len(arguments) == 2 {
		if !(arguments[1] == "shadow" || arguments[1] == "thinkertoy" || arguments[1] == "standard") {
			fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
			return false
		}
	}
	return true
}

func main() {
	arguments := os.Args[1:]

	if !validator(arguments) {
		return
	}
	var sentence string
	var banner string
	if len(arguments) == 1 {
		sentence = arguments[0]
		banner = ""
	}
	if len(arguments) == 2 {
		sentence = arguments[0]
		banner = arguments[1]
	}

	// Generate ASCII art using the Runner function
	result := Runner(sentence, banner)
	

	fmt.Println(result)
}
