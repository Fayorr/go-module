package main

import (
	"fmt"
	"os"
)

func validator(arguments []string) bool {

	if len(arguments) < 1 {
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

	result := Runner(content, arguments)

	fmt.Print(result)

}
