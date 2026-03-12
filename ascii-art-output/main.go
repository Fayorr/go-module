package main

import (
	"fmt"
	"os"
	"strings"
)

func validator(arguments []string) bool {
	// expect exactly three arguments: output flag, sentence, banner name
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
	fileName := strings.TrimPrefix(flag, "--output=")
	finalSen := strings.ReplaceAll(sentence, "\\n", "\n")
	result := Runner(finalSen, banner)

	resultBytes := []byte(result)
	err := os.WriteFile(fileName, resultBytes, 0644)
	if err != nil {
		fmt.Println("error writing to file")
		os.Exit(1)
	}
}
