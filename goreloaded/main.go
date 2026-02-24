package main

import (
	// "bufio"
	"fmt"
	"goreloaded/formatter"
	"os"
	"strings"
)

func createFile(path, text string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(text)
	return nil
}

func readFile(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
func validator() bool {
	arguments := os.Args

	if len(arguments) != 3 {
		fmt.Println("Error: Please provide exactly two files. Usage: go run . <input.txt> <output.txt>")
		return false
	}

	// 2. Safely access the arguments now that we know they exist
	inputFile := arguments[1]
	outputFile := arguments[2]

	// 3. Ensure they actually END with .txt
	if !strings.HasSuffix(inputFile, ".txt") || !strings.HasSuffix(outputFile, ".txt") {
		fmt.Println("Error: Wrong naming. Both files must end with .txt")
		return false
	}

	return true
}

func main() {
	if !validator() {
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]
	sentence, _ := readFile(inputFile)

	resStr := formatter.Runner(sentence)

	createFile(outputFile, resStr)

	fmt.Println("Validation passed. Ready to process:", inputFile, outputFile)
}
