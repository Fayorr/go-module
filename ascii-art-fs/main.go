package main

import (
	"fmt"
	"os"
)

func validator(arguments []string) bool {
	// Check for exactly 3 arguments
	if len(arguments) != 2 {
		fmt.Println("Kindly pass in 2 arguments")
		return false
	}
	if arguments[0] == "" {
		return false
	}
	if !(arguments[1] == "shadow" || arguments[1] == "thinkertoy" || arguments[1] == "standard") {
		fmt.Println("kindly pass in arguments: <sentence> <banner>")
		return false
	}
	return true
}

func main() {
	arguments := os.Args[1:]

	if !validator(arguments) {
		return
	}
	sentence := arguments[0]
	banner := arguments[1]

	// Generate ASCII art using the Runner function
	result := Runner(sentence, banner)

	fmt.Println(result)
}
