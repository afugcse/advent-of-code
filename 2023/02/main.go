package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
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

func parseInput(input string) any {
	return nil
}

func part1(parsedData any) any {
	return nil
}

func part2(parsedData any) any {
	return nil
}
