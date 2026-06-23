package main

import (
	"fmt"
	"os"
	"strings"
)

func printUsage() {
	fmt.Print("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard\n")
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	var align string
	var nonFlagArgs []string

	for _, arg := range args {
		if strings.HasPrefix(arg, "--align=") {
			if align != "" {
				printUsage()
				os.Exit(0)
			}
			align = strings.TrimPrefix(arg, "--align=")
			if align != "center" && align != "left" && align != "right" && align != "justify" {
				printUsage()
				os.Exit(0)
			}
		} else if strings.HasPrefix(arg, "-") {
			printUsage()
			os.Exit(0)
		} else {
			nonFlagArgs = append(nonFlagArgs, arg)
		}
	}

	if align == "" {
		align = "left"
	}

	var text, bannerType string
	if len(nonFlagArgs) == 1 {
		text = nonFlagArgs[0]
		bannerType = "standard"
	} else if len(nonFlagArgs) == 2 {
		text = nonFlagArgs[0]
		bannerType = nonFlagArgs[1]
	} else {
		printUsage()
		os.Exit(0)
	}

	if text == "" {
		return
	}

	Run(align, text, bannerType)
}