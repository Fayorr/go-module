package main

import (
	"fmt"
	"os"
	"strings"
)

func validator(arguments []string) bool {

	if len(arguments) != 3 {
		return false
	}
	if arguments[0] != "" && strings.HasSuffix(arguments[0], ".txt"){
		return false
	}

	return true
}
func readFile() (string, error) {
	//change the banner file name here
	content, err := os.ReadFile("./banner/shadow.txt")

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

	flag := arguments[0]
	sentence := arguments[1]
	banner := arguments[2]

	result := Runner(flag, sentence, banner)
	fmt.Print(result)
}
