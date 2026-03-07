package main

import (
	"fmt"
	"os"
)

// func colors() {
// 	color := map[string]string{
// 		"reset":   "\033[0m",
// 		"red":     "\033[31m",
// 		"green":   "\033[32m",
// 		"yellow":  "\033[33m",
// 		"blue":    "\033[34m",
// 		"magenta": "\033[35m",
// 		"cyan":    "\033[36m",
// 		"gray":    "\033[37m",
// 		"white":   "\033[97m",
// 	}

// }
func validator(arguments []string) bool {

	if len(arguments) < 1 || len(arguments) > 3 {
		return false
	}
	if arguments[0] == "" {
		return false
	}

	return true
}
func readFile() (string, error) {
	content, err := os.ReadFile("standard.txt")

	if err != nil {
		fmt.Println("Error in reading file")
		return "", err
	}

	return string(content), nil
}
func main() {
	arguments := os.Args[1:]

	if !validator(arguments) {
		return
	}

	content, _ := readFile()
	if len(arguments) == 2 {
		flag := arguments[0]
		sentence := arguments[1]
		subString := ""

		result := Runner(content, flag, subString, sentence)
		fmt.Print(result)
	}
	if len(arguments) == 3 {
		flag := arguments[0]
		subString := arguments[1]
		sentence := arguments[2]

		result := Runner(content, flag, subString, sentence)
		fmt.Print(result)
	}
}
