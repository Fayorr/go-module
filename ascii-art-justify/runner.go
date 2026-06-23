package main

import (
	"fmt"
	"os"
	"strings"

	"ascii-art-justify/internal"
)

// Run processes the parsed arguments and prints the aligned ASCII art
func Run(align, text, bannerType string) {
	bannerPath := "banner/" + bannerType + ".txt"
	
	banner, err := internal.LoadBanner(bannerPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	termWidth := internal.GetTermWidth()

	parts := strings.Split(text, "\\n")
	allEmpty := true
	for _, p := range parts {
		if p != "" {
			allEmpty = false
			break
		}
	}

	if allEmpty {
		for i := 0; i < len(parts)-1; i++ {
			fmt.Println()
		}
		return
	}

	for _, part := range parts {
		if part == "" {
			fmt.Println()
			continue
		}
		internal.AlignAndPrint(part, banner, align, termWidth)
	}
}