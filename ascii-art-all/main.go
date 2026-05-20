package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	arguments := os.Args[1:]

	if len(arguments) < 1 || len(arguments) > 3 {
		fmt.Println(`Usage: go run . [OPTION] [STRING] [BANNER]

EX: go run . --output=<fileName.txt> something standard`)
		return
	}
	var input string
	var banner string
	var fileName string


	if len(arguments) == 1 {
		input = arguments[0]
		banner = "standard"
	}
	if len(arguments) == 2 {
		if arguments[1] != "standard" && arguments[1] != "shadow" && arguments[1] != "thinkertoy" {
			fmt.Println(`Usage: go run . [STRING] [BANNER]

EX: go run . something standard`)
			return
		}
		input = arguments[0]
		banner = arguments[1]
	}
	if len(arguments) == 3 {
		if arguments[0] == "" {
			fmt.Println(`Usage: go run . [OPTION] [STRING] [BANNER]

EX: go run . --output=<fileName.txt> something standard`)
			return
		}
		if strings.HasPrefix(arguments[0], "--output=") {
			fileName = strings.TrimPrefix(arguments[0], "--output=")
		}
		input = arguments[1]
		banner = arguments[2]
	}
	if len(arguments) == 4 {
		input = arguments[0]
	}
	result := Runner(input, banner)

	if fileName != "" {
		os.WriteFile(fileName, []byte(result), 0644)
	}
	
}
