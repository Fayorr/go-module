package formatter

import (
	"strconv"
	"strings"
	"unicode"
)

func CorrectCase(sentence string) string {

	senArray := strings.Fields(sentence)
	for i := 0; i < len(senArray); i++ {
		if senArray[i] == "(hex)" && i-1 >= 0 {
			dec, _ := strconv.ParseInt(senArray[i-1], 16, 64)
			decStr := strconv.Itoa(int(dec))
			senArray[i-1] = decStr
			senArray = append(senArray[:i], senArray[i+1:]...)
			i--
		} else if senArray[i] == "(bin)" && i-1 >= 0 {
			dec, _ := strconv.ParseInt(senArray[i-1], 2, 64)
			decStr := strconv.Itoa(int(dec))
			senArray[i-1] = decStr
			senArray = append(senArray[:i], senArray[i+1:]...)
			i--
		} else if senArray[i] == "(low)" && i-1 >= 0 {
			senArray[i-1] = strings.ToLower(senArray[i-1])
			senArray = append(senArray[:i], senArray[i+1:]...)
			i--
		} else if senArray[i] == "(up)" && i-1 >= 0 {
			senArray[i-1] = strings.ToUpper(senArray[i-1])
			senArray = append(senArray[:i], senArray[i+1:]...)
			i--
		} else if senArray[i] == "(cap)" && i-1 >= 0 {
			wordSlice := []byte(senArray[i-1])
			wordSlice[0] = byte(unicode.ToUpper(rune(wordSlice[0])))
			senArray[i-1] = string(wordSlice)
			senArray = append(senArray[:i], senArray[i+1:]...)
			i--
		} else if (senArray[i] == "(low," || senArray[i] == "(cap," || senArray[i] == "(up,") && i+1 < len(senArray) {

			numStr := strings.TrimRight(senArray[i+1], ")")
			num, err := strconv.Atoi(numStr)

			if err == nil {
				for j := 1; j <= num; j++ {
					if i-j >= 0 {
						switch senArray[i] {
						case "(low,":
							senArray[i-j] = strings.ToLower(senArray[i-j])
						case "(up,":
							senArray[i-j] = strings.ToUpper(senArray[i-j])
						case "(cap,":
							wordSlice := []byte(senArray[i-j])
							wordSlice[0] = byte(unicode.ToUpper(rune(wordSlice[0])))
							senArray[i-j] = string(wordSlice)
						}

					}
				}
			}

			senArray = append(senArray[:i], senArray[i+2:]...)
			i--
		}
	}
	result := strings.Join(senArray, " ")
	return result
}
