package main

import (
	"fmt"
	"os"
	"strings"
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

	result, _ := readFile()
	
	
	lines := strings.Split(result, "\n")


	wordSlice := strings.Split(arguments[0], "\\n")

	for j:= 0; j<len(wordSlice); j++ {
		if wordSlice[j] == "" {
			fmt.Println()
			continue
		}
		for i:=1; i <= 8; i++ {
			for _, ch := range wordSlice[j] {
				pos := (int(ch - ' ')) * 9
				fmt.Print(lines[pos+i])
			}
			fmt.Println()
		}
	}
}
