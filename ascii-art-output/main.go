package main

import (
	"fmt"
	"os"
	"strings"
)

func validator(arguments []string) bool {
	// Check for exactly 3 arguments
	if len(arguments) != 3 {
		fmt.Println("need 3 arguments")
		return false
	}

	flag := arguments[0]
	if !strings.HasPrefix(flag, "--output=") {
		fmt.Println("needs the appropraite prefix")
		return false
	}
	return true
}
func main() {
	arguments := os.Args[1:]

	if !validator(arguments) {
		return
	}

	flag := arguments[0]
	sentence := arguments[1]
	banner := arguments[2]

	// Parse the output filename from the --output= flag
	fileName := strings.TrimPrefix(flag, "--output=")

	// Convert escaped newlines (\\n) to actual newlines for multi-line input
	finalSen := strings.ReplaceAll(sentence, "\\n", "\n")

	// Generate ASCII art using the Runner function
	result := Runner(finalSen, banner)

	// Convert string result to bytes and write to output file
	resultBytes := []byte(result)
	err := os.WriteFile(fileName, resultBytes, 0644)
	if err != nil {
		fmt.Println("error writing to file")
		os.Exit(1)
	}
}
