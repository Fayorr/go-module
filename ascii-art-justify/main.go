package main

import (
	"ascii-art-justify/internal"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func validator(arguments []string) bool {
	// Check for exactly 3 arguments
	if len(arguments) != 3 {
		fmt.Println("need 3 arguments")
		return false
	}

	flag := arguments[0]
	if !strings.HasPrefix(flag, "--align=") {
		fmt.Println("needs the appropraite prefix")
		return false
	}
	return true
}

func GetTerminalWidth() int {
   widthCmd := exec.Command("stty", "size")
    widthCmd.Stdin = os.Stdin 
    
    res, err := widthCmd.Output()
    if err != nil {
		fmt.Println("Can't find width output returning fallback width")
        return 80 
    }
	parts := strings.Fields(string(res))
	if len(parts) < 2 {
		return 80
	}
    width, err := strconv.Atoi(string(parts[1])) 
    if err != nil {
			fmt.Println("Can't convert width to int, returning fallback width")
        return 80 
    }
    return width
}


func main() {
	// arguments := os.Args[1:]

	// if !validator(arguments) {
	// 	return
	// }

	// flag := arguments[0]
	// sentence := arguments[1]
	// banner := arguments[2]

	// // Parse the alignment from the --align= flag
	// alignment := strings.TrimPrefix(flag, "--align=")

	// // Convert escaped newlines (\\n) to actual newlines for multi-line input
	// finalSen := strings.ReplaceAll(sentence, "\\n", "\n")

	// // Generate ASCII art using the Runner function
	// result := Runner(finalSen, banner)

	// // Convert string result to bytes and write to output file
	// fmt.Print(result)
	
	width := getTerminalWidth()
	internal.CalculatePadding(width int, side string, )
}
