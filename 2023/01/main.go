package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"
)

//go:embed input
var input string

func init() {
	// Remove trailing newline
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("No input file")
	}
}

func main() {
	pre, parsedData, post := time.Now(), parseInput(input), time.Now()
	fmt.Printf("Input parsed in %s\n", post.Sub(pre))

	pre, ans, post := time.Now(), part1(parsedData), time.Now()
	expectedAns := 54940
	fmt.Printf("Part1 answer: %d, in %s, pass: %t\n", ans, post.Sub(pre), ans == expectedAns)

	pre, ans, post = time.Now(), part2(parsedData), time.Now()
	expectedAns = 54208
	fmt.Printf("Part2 answer: %d, in %s, pass: %t\n", ans, post.Sub(pre), ans == expectedAns)
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

func part1(parsedData []string) int {
	sum := 0
	for _, line := range parsedData {
		firstNumber := ""
		secondNumber := ""

		i := 0
		j := len(line) - 1
		for i < j || firstNumber == "" || secondNumber == "" {
			if unicode.IsDigit(rune(line[i])) && firstNumber == "" {
				firstNumber = string(line[i])
			}

			if unicode.IsDigit(rune(line[j])) && secondNumber == "" {
				secondNumber = string(line[j])
			}

			i++
			j--
		}

		val, _ := strconv.Atoi(fmt.Sprintf("%s%s", firstNumber, secondNumber))
		sum += val
	}

	return sum
}

func part2(parsedData []string) int {
	sum := 0
	for _, line := range parsedData {
		firstNumber := ""
		secondNumber := ""

		i := 0
		j := len(line) - 1
		for i < j || firstNumber == "" || secondNumber == "" {
			wordDigit := getWordAsDigit(i, line, false)
			if (unicode.IsDigit(rune(line[i])) || wordDigit != "") && firstNumber == "" {
				if wordDigit != "" {
					firstNumber = wordDigit
				} else {
					firstNumber = string(line[i])
				}
			}

			wordDigit = getWordAsDigit(j, line, true)
			if (unicode.IsDigit(rune(line[j])) || wordDigit != "") && secondNumber == "" {
				if wordDigit != "" {
					secondNumber = wordDigit
				} else {
					secondNumber = string(line[j])
				}
			}

			i++
			j--
		}

		val, _ := strconv.Atoi(fmt.Sprintf("%s%s", firstNumber, secondNumber))
		sum += val
	}

	return sum
}

var wordToDigitMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func getWordAsDigit(idx int, line string, fromRight bool) string {
	if !fromRight {
		if idx+3 < len(line) {
			word := line[idx : idx+3]
			if val, ok := wordToDigitMap[word]; ok {
				return val
			}
		}

		if idx+4 < len(line) {
			word := line[idx : idx+4]
			if val, ok := wordToDigitMap[word]; ok {
				return val
			}
		}

		if idx+5 < len(line) {
			word := line[idx : idx+5]
			if val, ok := wordToDigitMap[word]; ok {
				return val
			}
		}
	} else {
		if idx-3+1 > 0 {
			word := line[idx-3+1 : idx+1]
			if val, ok := wordToDigitMap[word]; ok {
				return val
			}
		}

		if idx-4+1 > 0 {
			word := line[idx-4+1 : idx+1]
			if val, ok := wordToDigitMap[word]; ok {
				return val
			}
		}

		if idx-5+1 > 0 {
			word := line[idx-5+1 : idx+1]
			if val, ok := wordToDigitMap[word]; ok {
				return val
			}
		}
	}

	return ""
}
