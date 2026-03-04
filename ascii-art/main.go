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

	trimmedResult := strings.Split(result, "\n")

	word := strings.Join(arguments, " ")
	wordSlice := strings.Split(word, "\\n")

	fmt.Println(wordSlice)
	for j := 0; j < len(wordSlice); j++ {
		if wordSlice[j] == "" {
			continue
		}
		for i := 1; i <= 8; i++ {
			for _, ch := range wordSlice[j] {
				pos := (int(ch-' ') * 9)
				fmt.Print(trimmedResult[pos+i])
			}
			fmt.Println()
		}
	}

}
