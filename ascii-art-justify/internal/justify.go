package internal

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// LoadBanner reads the banner file and maps runes to 8 lines of text
func LoadBanner(filename string) (map[rune][]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("could not read %s", filename)
	}

	lines := strings.Split(string(content), "\n")
	for i := range lines {
		lines[i] = strings.TrimRight(lines[i], "\r")
	}

	banner := make(map[rune][]string)
	currentChar := rune(32)

	for i := 1; i < len(lines); i += 9 {
		if i+8 <= len(lines) {
			banner[currentChar] = lines[i : i+8]
			currentChar++
		}
	}
	return banner, nil
}

func GetTermWidth1() int {
	cmd := exec.Command("tput cols 2>dev/null | stty size 2>dev/null")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()

	if err == nil {
		w, err := strconv.Atoi(strings.TrimSpace(string(out)))
	}
}
// GetTermWidth dynamically determines the width of the terminal
func GetTermWidth() int {
	cmd := exec.Command("sh", "-c", "tput cols 2>/dev/null || stty size 2>/dev/null ")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err == nil {
		w, err := strconv.Atoi(strings.TrimSpace(string(out)))
		if err == nil && w > 0 {
			return w
		}
	}
	return 80 // Safe fallback
}

// AlignAndPrint calculates the layout and prints the output
func AlignAndPrint(text string, banner map[rune][]string, align string, termWidth int) {
	if align == "justify" {
		words := strings.Fields(text)
		if len(words) == 0 {
			fmt.Println()
			return
		}
		if len(words) == 1 {
			printLine(words[0], banner, "left", termWidth)
			return
		}

		var wordArts [][8]string
		totalWordsWidth := 0

		for _, w := range words {
			art := getArtForText(w, banner)
			wordArts = append(wordArts, art)
			totalWordsWidth += len(art[0])
		}

		totalSpaces := termWidth - totalWordsWidth
		if totalSpaces < 0 {
			totalSpaces = 0
		}

		baseGap := totalSpaces / (len(words) - 1)
		extraGap := totalSpaces % (len(words) - 1)

		for i := 0; i < 8; i++ {
			line := ""
			for j, wArt := range wordArts {
				line += wArt[i]
				if j < len(words)-1 {
					spaces := baseGap
					if j < extraGap {
						spaces++
					}
					line += strings.Repeat(" ", spaces)
				}
			}
			fmt.Println(line)
		}
	} else {
		printLine(text, banner, align, termWidth)
	}
}

func printLine(text string, banner map[rune][]string, align string, termWidth int) {
	art := getArtForText(text, banner)
	artWidth := len(art[0])

	padLeft := 0
	if align == "center" {
		padLeft = (termWidth - artWidth) / 2
	} else if align == "right" {
		padLeft = termWidth - artWidth
	}
	if padLeft < 0 {
		padLeft = 0
	}

	padding := strings.Repeat(" ", padLeft)
	for i := 0; i < 8; i++ {
		fmt.Println(padding + art[i])
	}
}

func getArtForText(text string, banner map[rune][]string) [8]string {
	var result [8]string
	for _, char := range text {
		if lines, ok := banner[char]; ok {
			for i := 0; i < 8; i++ {
				result[i] += lines[i]
			}
		}
	}
	return result
}